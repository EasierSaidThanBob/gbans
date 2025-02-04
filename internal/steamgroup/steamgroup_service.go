package steamgroup

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leighmacdonald/gbans/internal/domain"
	"github.com/leighmacdonald/gbans/internal/httphelper"
	"github.com/leighmacdonald/gbans/pkg/log"
	"github.com/leighmacdonald/gbans/pkg/util"
)

type steamgroupHandler struct {
	bgu domain.BanGroupUsecase
}

func NewSteamgroupHandler(engine *gin.Engine, bgu domain.BanGroupUsecase, ath domain.AuthUsecase) {
	handler := steamgroupHandler{
		bgu: bgu,
	}

	// mod
	modGrp := engine.Group("/")
	{
		mod := modGrp.Use(ath.AuthMiddleware(domain.PUser))
		mod.POST("/api/bans/group/create", handler.onAPIPostBansGroupCreate())
		mod.GET("/api/bans/group", handler.onAPIGetBansGroup())
		mod.DELETE("/api/bans/group/:ban_group_id", handler.onAPIDeleteBansGroup())
		mod.POST("/api/bans/group/:ban_group_id", handler.onAPIPostBansGroupUpdate())
	}
}

func (h steamgroupHandler) onAPIPostBansGroupCreate() gin.HandlerFunc {
	type apiBanGroupRequest struct {
		domain.TargetIDField
		domain.TargetGIDField
		Duration   string    `json:"duration"`
		Note       string    `json:"note"`
		ValidUntil time.Time `json:"valid_until"`
	}

	return func(ctx *gin.Context) {
		var req apiBanGroupRequest
		if !httphelper.Bind(ctx, &req) {
			return
		}

		groupID, errGroupID := req.TargetGroupID(ctx)
		if !errGroupID {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrBadRequest)

			return
		}

		var existing domain.BanGroup
		if errExist := h.bgu.GetByGID(ctx, groupID, &existing); errExist != nil {
			if !errors.Is(errExist, domain.ErrNoResult) {
				httphelper.ResponseErr(ctx, http.StatusConflict, domain.ErrDuplicate)

				return
			}
		}

		var (
			banSteamGroup domain.BanGroup
			sid           = httphelper.CurrentUserProfile(ctx).SteamID
		)

		duration, errDuration := util.CalcDuration(req.Duration, req.ValidUntil)
		if errDuration != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrBadRequest)

			return
		}

		targetID, targetIDOk := req.TargetSteamID(ctx)
		if !targetIDOk {
			httphelper.ErrorHandled(ctx, domain.ErrTargetID)

			return
		}

		groupID, groupIDOk := req.TargetGroupID(ctx)
		if !groupIDOk {
			httphelper.ErrorHandled(ctx, domain.ErrTargetID)

			return
		}

		if errBanSteamGroup := domain.NewBanSteamGroup(sid, targetID, duration, req.Note, domain.Web, groupID,
			"", domain.Banned, &banSteamGroup); errBanSteamGroup != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrBadRequest)
			slog.Error("Failed to save group ban", log.ErrAttr(errBanSteamGroup))

			return
		}

		if errBan := h.bgu.Ban(ctx, &banSteamGroup); errBan != nil {
			if errors.Is(errBan, domain.ErrDuplicate) {
				httphelper.ResponseErr(ctx, http.StatusConflict, domain.ErrDuplicate)

				return
			}

			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		ctx.JSON(http.StatusCreated, banSteamGroup)
	}
}

func (h steamgroupHandler) onAPIGetBansGroup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.GroupBansQueryFilter
		if !httphelper.BindQuery(ctx, &req) {
			return
		}

		banGroups, errBans := h.bgu.Get(ctx, req)
		if errBans != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)
			slog.Error("Failed to fetch banGroups", log.ErrAttr(errBans))

			return
		}

		ctx.JSON(http.StatusOK, banGroups)
	}
}

func (h steamgroupHandler) onAPIDeleteBansGroup() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		groupID, groupIDErr := httphelper.GetInt64Param(ctx, "ban_group_id")
		if groupIDErr != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInternal)

			return
		}

		var req domain.UnbanRequest
		if !httphelper.Bind(ctx, &req) {
			return
		}

		var banGroup domain.BanGroup
		if errFetch := h.bgu.GetByID(ctx, groupID, &banGroup); errFetch != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInternal)

			return
		}

		banGroup.UnbanReasonText = req.UnbanReasonText
		banGroup.Deleted = true

		if errSave := h.bgu.Save(ctx, &banGroup); errSave != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)
			slog.Error("Failed to delete asn ban", log.ErrAttr(errSave))

			return
		}

		banGroup.BanGroupID = 0

		ctx.JSON(http.StatusOK, banGroup)
	}
}

func (h steamgroupHandler) onAPIPostBansGroupUpdate() gin.HandlerFunc {
	type apiBanUpdateRequest struct {
		domain.TargetIDField
		Note       string    `json:"note"`
		ValidUntil time.Time `json:"valid_until"`
	}

	return func(ctx *gin.Context) {
		banGroupID, banIDErr := httphelper.GetInt64Param(ctx, "ban_group_id")
		if banIDErr != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInvalidParameter)

			return
		}

		var req apiBanUpdateRequest
		if !httphelper.Bind(ctx, &req) {
			return
		}

		targetSID, sidValid := req.TargetSteamID(ctx)
		if !sidValid {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInvalidParameter)

			return
		}

		var ban domain.BanGroup

		if errExist := h.bgu.GetByID(ctx, banGroupID, &ban); errExist != nil {
			if !errors.Is(errExist, domain.ErrNoResult) {
				httphelper.ResponseErr(ctx, http.StatusNotFound, domain.ErrNotFound)

				return
			}

			httphelper.ResponseErr(ctx, http.StatusConflict, domain.ErrDuplicate)

			return
		}

		ban.Note = req.Note
		ban.ValidUntil = req.ValidUntil
		ban.TargetID = targetSID

		if errSave := h.bgu.Save(ctx, &ban); errSave != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		ctx.JSON(http.StatusOK, ban)
	}
}
