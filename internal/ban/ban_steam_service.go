package ban

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leighmacdonald/gbans/internal/domain"
	"github.com/leighmacdonald/gbans/internal/httphelper"
	"github.com/leighmacdonald/gbans/internal/thirdparty"
	"github.com/leighmacdonald/gbans/pkg/util"
	"github.com/leighmacdonald/steamid/v3/steamid"
	"go.uber.org/zap"
)

type nanHandler struct {
	du  domain.DiscordUsecase
	bu  domain.BanSteamUsecase
	pu  domain.PersonUsecase
	cu  domain.ConfigUsecase
	log *zap.Logger
}

func NewBanHandler(logger *zap.Logger, engine *gin.Engine, bu domain.BanSteamUsecase, du domain.DiscordUsecase,
	pu domain.PersonUsecase, cu domain.ConfigUsecase, ath domain.AuthUsecase,
) {
	handler := nanHandler{log: logger, bu: bu, du: du, pu: pu, cu: cu}

	engine.GET("/api/stats", handler.onAPIGetStats())
	engine.GET("/export/bans/tf2bd", handler.onAPIExportBansTF2BD())
	engine.GET("/export/bans/valve/steamid", handler.onAPIExportBansValveSteamID())

	// auth
	authedGrp := engine.Group("/")
	{
		authed := authedGrp.Use(ath.AuthMiddleware(domain.PUser))
		authed.GET("/api/bans/steam/:ban_id", handler.onAPIGetBanByID())
		authed.GET("/api/sourcebans/:steam_id", handler.onAPIGetSourceBans())
	}

	// mod
	modGrp := engine.Group("/")
	{
		mod := modGrp.Use(ath.AuthMiddleware(domain.PModerator))

		mod.POST("/api/bans/steam", handler.onAPIGetBansSteam())
		mod.POST("/api/bans/steam/create", handler.onAPIPostBanSteamCreate())
		mod.DELETE("/api/bans/steam/:ban_id", handler.onAPIPostBanDelete())
		mod.POST("/api/bans/steam/:ban_id", handler.onAPIPostBanUpdate())
		mod.POST("/api/bans/steam/:ban_id/status", handler.onAPIPostSetBanAppealStatus())
	}
}

func (h nanHandler) onAPIPostSetBanAppealStatus() gin.HandlerFunc {
	type setStatusReq struct {
		AppealState domain.AppealState `json:"appeal_state"`
	}

	log := h.log.Named(runtime.FuncForPC(make([]uintptr, 10)[0]).Name())

	return func(ctx *gin.Context) {
		banID, banIDErr := httphelper.GetInt64Param(ctx, "ban_id")
		if banIDErr != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInvalidParameter)

			return
		}

		var req setStatusReq
		if !httphelper.Bind(ctx, log, &req) {
			return
		}

		bannedPerson := domain.NewBannedPerson()
		if banErr := h.bu.GetByBanID(ctx, banID, &bannedPerson, false); banErr != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		if bannedPerson.AppealState == req.AppealState {
			httphelper.ResponseErr(ctx, http.StatusConflict, domain.ErrStateUnchanged)

			return
		}

		original := bannedPerson.AppealState
		bannedPerson.AppealState = req.AppealState

		if errSave := h.bu.Save(ctx, &bannedPerson.BanSteam); errSave != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{})

		log.Info("Updated ban appeal state",
			zap.Int64("ban_id", banID),
			zap.Int("from_state", int(original)),
			zap.Int("to_state", int(req.AppealState)))
	}
}

type apiBanRequest struct {
	SourceID       domain.StringSID `json:"source_id"`
	TargetID       domain.StringSID `json:"target_id"`
	Duration       string           `json:"duration"`
	ValidUntil     time.Time        `json:"valid_until"`
	BanType        domain.BanType   `json:"ban_type"`
	Reason         domain.Reason    `json:"reason"`
	ReasonText     string           `json:"reason_text"`
	Note           string           `json:"note"`
	ReportID       int64            `json:"report_id"`
	DemoName       string           `json:"demo_name"`
	DemoTick       int              `json:"demo_tick"`
	IncludeFriends bool             `json:"include_friends"`
}

func (h nanHandler) onAPIPostBanSteamCreate() gin.HandlerFunc {
	log := h.log.Named(runtime.FuncForPC(make([]uintptr, 10)[0]).Name())

	return func(ctx *gin.Context) {
		var req apiBanRequest
		if !httphelper.Bind(ctx, log, &req) {
			return
		}

		var (
			origin   = domain.Web
			sid      = httphelper.CurrentUserProfile(ctx).SteamID
			sourceID = domain.StringSID(sid.String())
		)

		// srcds sourced bans provide a source_id to id the admin
		if req.SourceID != "" {
			sourceID = req.SourceID
			origin = domain.InGame
		}

		duration, errDuration := util.CalcDuration(req.Duration, req.ValidUntil)
		if errDuration != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrBadRequest)

			return
		}

		var banSteam domain.BanSteam
		if errBanSteam := domain.NewBanSteam(ctx,
			sourceID,
			req.TargetID,
			duration,
			req.Reason,
			req.ReasonText,
			req.Note,
			origin,
			req.ReportID,
			req.BanType,
			req.IncludeFriends,
			&banSteam,
		); errBanSteam != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrBadRequest)

			return
		}

		if errBan := h.bu.Ban(ctx, &banSteam); errBan != nil {
			log.Error("Failed to ban steam profile",
				zap.Error(errBan), zap.Int64("target_id", banSteam.TargetID.Int64()))

			if errors.Is(errBan, domain.ErrDuplicate) {
				httphelper.ResponseErr(ctx, http.StatusConflict, domain.ErrDuplicate)

				return
			}

			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)
			log.Error("Failed to save new steam ban", zap.Error(errBan))

			return
		}

		ctx.JSON(http.StatusCreated, banSteam)
	}
}

func (h nanHandler) onAPIGetBanByID() gin.HandlerFunc {
	log := h.log.Named(runtime.FuncForPC(make([]uintptr, 10)[0]).Name())

	return func(ctx *gin.Context) {
		curUser := httphelper.CurrentUserProfile(ctx)

		banID, errID := httphelper.GetInt64Param(ctx, "ban_id")
		if errID != nil || banID == 0 {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInvalidParameter)

			return
		}

		deletedOk := false

		fullValue, fullOk := ctx.GetQuery("deleted")
		if fullOk {
			deleted, deletedOkErr := strconv.ParseBool(fullValue)
			if deletedOkErr != nil {
				log.Error("Failed to parse ban full query value", zap.Error(deletedOkErr))
			} else {
				deletedOk = deleted
			}
		}

		bannedPerson := domain.NewBannedPerson()
		if errGetBan := h.bu.GetByBanID(ctx, banID, &bannedPerson, deletedOk); errGetBan != nil {
			httphelper.ResponseErr(ctx, http.StatusNotFound, domain.ErrNotFound)
			log.Error("Failed to fetch steam ban", zap.Error(errGetBan))

			return
		}

		if !httphelper.CheckPrivilege(ctx, curUser, steamid.Collection{bannedPerson.TargetID}, domain.PModerator) {
			return
		}

		ctx.JSON(http.StatusOK, bannedPerson)
	}
}

func (h nanHandler) onAPIGetSourceBans() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		steamID, errID := httphelper.GetSID64Param(ctx, "steam_id")
		if errID != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInvalidParameter)

			return
		}

		records, errRecords := thirdparty.BDSourceBans(ctx, steamID)
		if errRecords != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		ctx.JSON(http.StatusOK, records)
	}
}

func (h nanHandler) onAPIGetStats() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var stats domain.Stats
		if errGetStats := h.bu.Stats(ctx, &stats); errGetStats != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		stats.ServersAlive = 1

		ctx.JSON(http.StatusOK, stats)
	}
}

func (h nanHandler) onAPIExportBansValveSteamID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bans, _, errBans := h.bu.Get(ctx, domain.SteamBansQueryFilter{
			BansQueryFilter: domain.BansQueryFilter{PermanentOnly: true},
		})

		if errBans != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		var entries []string

		for _, ban := range bans {
			if ban.Deleted ||
				!ban.IsEnabled {
				continue
			}

			entries = append(entries, fmt.Sprintf("banid 0 %s", steamid.SID64ToSID(ban.TargetID)))
		}

		ctx.Data(http.StatusOK, "text/plain", []byte(strings.Join(entries, "\n")))
	}
}

func (h nanHandler) onAPIExportBansTF2BD() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO limit / make specialized query since this returns all results
		bans, _, errBans := h.bu.Get(ctx, domain.SteamBansQueryFilter{
			BansQueryFilter: domain.BansQueryFilter{
				QueryFilter: domain.QueryFilter{
					Deleted: false,
				},
			},
		})

		if errBans != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		var filtered []domain.BannedSteamPerson

		for _, ban := range bans {
			if ban.Reason != domain.Cheating ||
				ban.Deleted ||
				!ban.IsEnabled {
				continue
			}

			filtered = append(filtered, ban)
		}

		config := h.cu.Config()

		out := thirdparty.TF2BDSchema{
			Schema: "https://raw.githubusercontent.com/PazerOP/tf2_bot_detector/master/schemas/v3/playerlist.schema.json",
			FileInfo: thirdparty.FileInfo{
				Authors:     []string{config.General.SiteName},
				Description: "Players permanently banned for cheating",
				Title:       fmt.Sprintf("%s Cheater List", config.General.SiteName),
				UpdateURL:   h.cu.ExtURLRaw("/export/bans/tf2bd"),
			},
			Players: []thirdparty.Players{},
		}

		for _, ban := range filtered {
			out.Players = append(out.Players, thirdparty.Players{
				Attributes: []string{"cheater"},
				Steamid:    ban.TargetID,
				LastSeen: thirdparty.LastSeen{
					PlayerName: ban.TargetPersonaname,
					Time:       int(ban.UpdatedOn.Unix()),
				},
			})
		}

		ctx.JSON(http.StatusOK, out)
	}
}

func (h nanHandler) onAPIGetBansSteam() gin.HandlerFunc {
	log := h.log.Named(runtime.FuncForPC(make([]uintptr, 10)[0]).Name())

	return func(ctx *gin.Context) {
		var req domain.SteamBansQueryFilter
		if !httphelper.Bind(ctx, log, &req) {
			return
		}

		bans, count, errBans := h.bu.Get(ctx, req)
		if errBans != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)
			log.Error("Failed to fetch steam bans", zap.Error(errBans))

			return
		}

		ctx.JSON(http.StatusOK, domain.NewLazyResult(count, bans))
	}
}

func (h nanHandler) onAPIPostBanDelete() gin.HandlerFunc {
	log := h.log.Named(runtime.FuncForPC(make([]uintptr, 10)[0]).Name())

	return func(ctx *gin.Context) {
		banID, banIDErr := httphelper.GetInt64Param(ctx, "ban_id")
		if banIDErr != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInvalidParameter)

			return
		}

		var req domain.UnbanRequest
		if !httphelper.Bind(ctx, log, &req) {
			return
		}

		bannedPerson := domain.NewBannedPerson()
		if banErr := h.bu.GetByBanID(ctx, banID, &bannedPerson, false); banErr != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		changed, errSave := h.bu.Unban(ctx, bannedPerson.TargetID, req.UnbanReasonText)
		if errSave != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		if !changed {
			httphelper.ResponseErr(ctx, http.StatusNotFound, domain.ErrUnbanFailed)

			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{})
	}
}

func (h nanHandler) onAPIPostBanUpdate() gin.HandlerFunc {
	log := h.log.Named(runtime.FuncForPC(make([]uintptr, 10)[0]).Name())

	type updateBanRequest struct {
		TargetID       domain.StringSID `json:"target_id"`
		BanType        domain.BanType   `json:"ban_type"`
		Reason         domain.Reason    `json:"reason"`
		ReasonText     string           `json:"reason_text"`
		Note           string           `json:"note"`
		IncludeFriends bool             `json:"include_friends"`
		ValidUntil     time.Time        `json:"valid_until"`
	}

	return func(ctx *gin.Context) {
		banID, banIDErr := httphelper.GetInt64Param(ctx, "ban_id")
		if banIDErr != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInvalidParameter)

			return
		}

		var req updateBanRequest
		if !httphelper.Bind(ctx, log, &req) {
			return
		}

		if time.Since(req.ValidUntil) > 0 {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrBadRequest)

			return
		}

		bannedPerson := domain.NewBannedPerson()
		if banErr := h.bu.GetByBanID(ctx, banID, &bannedPerson, false); banErr != nil {
			httphelper.ResponseErr(ctx, http.StatusNotFound, domain.ErrNotFound)

			return
		}

		if req.Reason == domain.Custom {
			if req.ReasonText == "" {
				httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrBadRequest)

				return
			}

			bannedPerson.ReasonText = req.ReasonText
		} else {
			bannedPerson.ReasonText = ""
		}

		bannedPerson.Note = req.Note
		bannedPerson.BanType = req.BanType
		bannedPerson.Reason = req.Reason
		bannedPerson.IncludeFriends = req.IncludeFriends
		bannedPerson.ValidUntil = req.ValidUntil

		if errSave := h.bu.Save(ctx, &bannedPerson.BanSteam); errSave != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)
			log.Error("Failed to save updated ban", zap.Error(errSave))

			return
		}

		ctx.JSON(http.StatusAccepted, bannedPerson)
	}
}
