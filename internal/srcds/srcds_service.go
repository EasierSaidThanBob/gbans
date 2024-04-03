package srcds

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/netip"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/leighmacdonald/gbans/internal/discord"
	"github.com/leighmacdonald/gbans/internal/domain"
	"github.com/leighmacdonald/gbans/internal/httphelper"
	"github.com/leighmacdonald/gbans/internal/thirdparty"
	"github.com/leighmacdonald/gbans/pkg/log"
	"github.com/leighmacdonald/gbans/pkg/util"
	"github.com/leighmacdonald/steamid/v4/steamid"
)

type srcdsHandler struct {
	sru             domain.SRCDSUsecase
	ServerUsecase   domain.ServersUsecase
	PersonUsecase   domain.PersonUsecase
	stateUsecase    domain.StateUsecase
	discordUsecase  domain.DiscordUsecase
	configUsecase   domain.ConfigUsecase
	reportUsecase   domain.ReportUsecase
	assetUsecase    domain.AssetUsecase
	banUsecase      domain.BanSteamUsecase
	banGroupUsecase domain.BanGroupUsecase
	banASNUsecase   domain.BanASNUsecase
	banNetUsecase   domain.BanNetUsecase
	networkUsecase  domain.NetworkUsecase
	demoUsecase     domain.DemoUsecase
}

const authTokenDuration = time.Minute * 15

func NewSRCDSHandler(engine *gin.Engine, srcdsUsecase domain.SRCDSUsecase, serversUsecase domain.ServersUsecase,
	personUsecase domain.PersonUsecase, assetUsecase domain.AssetUsecase, reportUsecase domain.ReportUsecase,
	banUsecase domain.BanSteamUsecase, networkUsecase domain.NetworkUsecase, banGroupUsecase domain.BanGroupUsecase,
	demoUsecase domain.DemoUsecase, authUsecase domain.AuthUsecase, banASNUsecase domain.BanASNUsecase, banNetUsecase domain.BanNetUsecase,
	configUsecase domain.ConfigUsecase, discordUsecase domain.DiscordUsecase, stateUsecase domain.StateUsecase,
) {
	handler := srcdsHandler{
		sru:             srcdsUsecase,
		ServerUsecase:   serversUsecase,
		PersonUsecase:   personUsecase,
		reportUsecase:   reportUsecase,
		banUsecase:      banUsecase,
		assetUsecase:    assetUsecase,
		networkUsecase:  networkUsecase,
		banGroupUsecase: banGroupUsecase,
		demoUsecase:     demoUsecase,
		banASNUsecase:   banASNUsecase,
		banNetUsecase:   banNetUsecase,
		configUsecase:   configUsecase,
		discordUsecase:  discordUsecase,
		stateUsecase:    stateUsecase,
	}

	// unauthed
	engine.POST("/api/server/auth", handler.onSAPIPostServerAuth())

	// serverAuth := srvGrp.Use(authServerMiddleWare(env))
	srvGrp := engine.Group("/")
	{
		server := srvGrp.Use(authUsecase.AuthServerMiddleWare())
		server.GET("/api/server/admins", handler.onAPIGetServerAdmins())
		server.POST("/api/ping_mod", handler.onAPIPostPingMod())
		server.POST("/api/check", handler.onAPIPostServerCheck())
		server.POST("/api/demo", handler.onAPIPostDemo())

		// Duplicated since we need to authenticate via server middleware
		server.POST("/api/sm/bans/steam/create", handler.onAPIPostBanSteamCreate())
		server.POST("/api/sm/report/create", handler.onAPIPostReportCreate())
		server.POST("/api/state_update", handler.onAPIPostServerState())
	}
}

type ServerAuthResp struct {
	Status bool   `json:"status"`
	Token  string `json:"token"`
}

func newServerToken(serverID int, cookieKey string) (string, error) {
	curTime := time.Now()

	claims := &domain.ServerAuthClaims{
		ServerID: serverID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(curTime.Add(authTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(curTime),
			NotBefore: jwt.NewNumericDate(curTime),
		},
	}

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, errSigned := tokenWithClaims.SignedString([]byte(cookieKey))
	if errSigned != nil {
		return "", errors.Join(errSigned, domain.ErrSignJWT)
	}

	return signedToken, nil
}

func (s *srcdsHandler) onSAPIPostServerAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.ServerAuthReq
		if !httphelper.Bind(ctx, &req) {
			return
		}

		token, err := s.sru.ServerAuth(ctx, req)
		if err != nil {
			httphelper.ResponseErr(ctx, http.StatusUnauthorized, domain.ErrPermissionDenied)
			slog.Warn("Failed to check server auth", log.ErrAttr(err))

			return
		}

		ctx.JSON(http.StatusOK, ServerAuthResp{Status: true, Token: token})
	}
}

func (s *srcdsHandler) onAPIPostServerState() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.PartialStateUpdate
		if !httphelper.Bind(ctx, &req) {
			return
		}

		serverID := httphelper.ServerFromCtx(ctx) // TODO use generic func for int
		if serverID == 0 {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrParamInvalid)

			return
		}

		if errUpdate := s.stateUsecase.Update(serverID, req); errUpdate != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		ctx.AbortWithStatus(http.StatusNoContent)
	}
}

func (s *srcdsHandler) onAPIPostReportCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		currentUser := httphelper.CurrentUserProfile(ctx)

		var req domain.CreateReportReq
		if !httphelper.Bind(ctx, &req) {
			return
		}

		report, errReport := s.sru.Report(ctx, currentUser, req)
		if errReport != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, errReport)
			slog.Error("Failed to create report", log.ErrAttr(errReport))

			return
		}

		ctx.JSON(http.StatusCreated, report)
	}
}

func (s *srcdsHandler) onAPIPostDemo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		demoFormFile, errDemoFile := ctx.FormFile("demo")
		if errDemoFile != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrBadRequest)

			return
		}

		demoFile, errDemoHandle := demoFormFile.Open()
		if errDemoHandle != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrBadRequest)

			return
		}

		defer util.LogCloser(demoFile)

		newDemo, errCreateDemo := s.demoUsecase.Create(ctx, demoFormFile.Filename, demoFile, demoFormFile.Filename, httphelper.ServerFromCtx(ctx))
		if errCreateDemo != nil {
			httphelper.HandleErrInternal(ctx)

			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"demo_id": newDemo.DemoID})
	}
}

type apiBanRequest struct {
	SourceID       steamid.SteamID `json:"source_id"`
	TargetID       steamid.SteamID `json:"target_id"`
	Duration       string          `json:"duration"`
	ValidUntil     time.Time       `json:"valid_until"`
	BanType        domain.BanType  `json:"ban_type"`
	Reason         domain.Reason   `json:"reason"`
	ReasonText     string          `json:"reason_text"`
	Note           string          `json:"note"`
	ReportID       int64           `json:"report_id"`
	DemoName       string          `json:"demo_name"`
	DemoTick       int             `json:"demo_tick"`
	IncludeFriends bool            `json:"include_friends"`
}

func (s *srcdsHandler) onAPIPostBanSteamCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req apiBanRequest
		if !httphelper.Bind(ctx, &req) {
			return
		}

		var (
			origin   = domain.Web
			curUser  = httphelper.CurrentUserProfile(ctx)
			sourceID = curUser.SteamID
		)

		// srcds sourced bans provide a source_id to id the admin
		if req.SourceID.Valid() {
			sourceID = req.SourceID
			origin = domain.InGame
		}

		duration, errDuration := util.CalcDuration(req.Duration, req.ValidUntil)
		if errDuration != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrBadRequest)

			return
		}

		var banSteam domain.BanSteam
		if errBanSteam := domain.NewBanSteam(sourceID, req.TargetID, duration, req.Reason, req.ReasonText, req.Note, origin,
			req.ReportID, req.BanType, req.IncludeFriends, &banSteam); errBanSteam != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrBadRequest)

			return
		}

		if errBan := s.banUsecase.Ban(ctx, curUser, &banSteam); errBan != nil {
			slog.Error("Failed to ban steam profile",
				log.ErrAttr(errBan), slog.Int64("target_id", banSteam.TargetID.Int64()))

			if errors.Is(errBan, domain.ErrDuplicate) {
				httphelper.ResponseErr(ctx, http.StatusConflict, domain.ErrDuplicate)

				return
			}

			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)
			slog.Error("Failed to save new steam ban", log.ErrAttr(errBan))

			return
		}

		ctx.JSON(http.StatusCreated, banSteam)
	}
}

func (s *srcdsHandler) onAPIGetServerAdmins() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		perms, err := s.ServerUsecase.GetServerPermissions(ctx)
		if err != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		ctx.JSON(http.StatusOK, perms)
	}
}

type pingReq struct {
	ServerName string          `json:"server_name"`
	Name       string          `json:"name"`
	SteamID    steamid.SteamID `json:"steam_id"`
	Reason     string          `json:"reason"`
	Client     int             `json:"client"`
}

func (s *srcdsHandler) onAPIPostPingMod() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req pingReq
		if !httphelper.Bind(ctx, &req) {
			return
		}

		conf := s.configUsecase.Config()
		players := s.stateUsecase.Find("", req.SteamID, nil, nil)

		if len(players) == 0 && conf.General.Mode != domain.TestMode {
			slog.Error("Failed to find player on /mod call")
			httphelper.ResponseErr(ctx, http.StatusFailedDependency, domain.ErrInternal)

			return
		}

		ctx.JSON(http.StatusOK, gin.H{"client": req.Client, "message": "Moderators have been notified"})

		if !conf.Discord.Enabled {
			return
		}

		author, err := s.PersonUsecase.GetOrCreatePersonBySteamID(ctx, req.SteamID)
		if err != nil {
			slog.Error("Failed to load user", log.ErrAttr(err))

			return
		}

		s.discordUsecase.SendPayload(domain.ChannelMod,
			discord.PingModMessage(author, conf.ExtURL(author), req.Reason, req.ServerName, conf.Discord.ModPingRoleID))
	}
}

type CheckRequest struct {
	ClientID int        `json:"client_id"`
	SteamID  string     `json:"steam_id"`
	IP       netip.Addr `json:"ip"`
	Name     string     `json:"name,omitempty"`
}

type CheckResponse struct {
	ClientID        int              `json:"client_id"`
	SteamID         steamid.SID      `json:"steam_id"`
	BanType         domain.BanType   `json:"ban_type"`
	PermissionLevel domain.Privilege `json:"permission_level"`
	Msg             string           `json:"msg"`
}

// onAPIPostServerCheck takes care of checking if the player connecting to the server is
// allowed to connect, or otherwise has restrictions such as being mutes. It performs
// the following actions/checks in order:
//
// - Add ip to connection history
// - Check if is a part of a steam group ban
// - Check if ip belongs to banned 3rd party CIDR block, like VPNs.
// - Check if ip belongs to one or more local CIDR bans
// - Check if ip belongs to a banned AS Number range
// - Check if steam_id is part of a local steam ban
// - Check if player is connecting from an IP that belongs to a banned player
//
// Returns an ok/muted/banned status for the player.
func (s *srcdsHandler) onAPIPostServerCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request CheckRequest
		if errBind := ctx.BindJSON(&request); errBind != nil { // we don't currently use bind() for server api
			ctx.JSON(http.StatusInternalServerError, CheckResponse{
				BanType: domain.Unknown,
				Msg:     "Error determining state",
			})

			return
		}

		sid := steamid.New(request.SteamID)

		slog.Debug("Player connecting",
			slog.String("ip", request.IP.String()),
			slog.Int64("sid64", sid.Int64()),
			slog.String("name", request.Name))

		resp := CheckResponse{ClientID: request.ClientID, SteamID: sid.Steam(false), BanType: domain.Unknown, Msg: ""}

		responseCtx, cancelResponse := context.WithTimeout(ctx, time.Second*15)
		defer cancelResponse()

		if !sid.Valid() {
			resp.Msg = "Invalid steam id"
			ctx.JSON(http.StatusBadRequest, resp)

			return
		}

		person, errPerson := s.PersonUsecase.GetOrCreatePersonBySteamID(responseCtx, sid)
		if errPerson != nil {
			slog.Error("Failed to create connecting player", log.ErrAttr(errPerson))
		} else if person.Expired() {
			if err := thirdparty.UpdatePlayerSummary(ctx, &person); err != nil {
				slog.Error("Failed to update connecting player", log.ErrAttr(err))
			} else {
				if errSave := s.PersonUsecase.SavePerson(ctx, &person); errSave != nil {
					slog.Error("Failed to save connecting player summary", log.ErrAttr(err))
				}
			}
		}

		if errAddHist := s.networkUsecase.AddConnectionHistory(ctx, &domain.PersonConnection{
			IPAddr:      request.IP,
			SteamID:     sid,
			PersonaName: request.Name,
			CreatedOn:   time.Now(),
			ServerID:    ctx.GetInt("server_id"),
		}); errAddHist != nil {
			slog.Error("Failed to add conn history", log.ErrAttr(errAddHist))
		}

		resp.PermissionLevel = person.PermissionLevel

		if s.checkGroupBan(ctx, sid, &resp) || s.checkFriendBan(ctx, sid, &resp) {
			return
		}

		if s.checkNetBlockBan(ctx, sid, request.IP, &resp) {
			return
		}

		if s.checkIPBan(ctx, sid, request.IP, responseCtx, &resp) {
			return
		}

		if s.checkASN(ctx, sid, request.IP, responseCtx, &resp) {
			return
		}

		bannedPerson, errGetBan := s.banUsecase.GetBySteamID(responseCtx, sid, false)
		if errGetBan != nil {
			if errors.Is(errGetBan, domain.ErrNoResult) {
				isShared, errShared := s.banUsecase.IsOnIPWithBan(ctx, httphelper.CurrentUserProfile(ctx),
					sid, request.IP)
				if errShared != nil {
					slog.Error("Failed to check shared ip state", log.ErrAttr(errShared))

					ctx.JSON(http.StatusOK, resp)

					return
				}

				if isShared {
					slog.Info("Player connected from IP of a banned player",
						slog.String("steam_id", sid.String()),
						slog.String("ip", request.IP.String()))

					resp.BanType = domain.Banned
					resp.Msg = "Ban evasion. Previous ban updated to permanent if not already permanent"

					ctx.JSON(http.StatusOK, resp)

					return
				}

				// No ban, exit early
				resp.BanType = domain.OK
				ctx.JSON(http.StatusOK, resp)

				return
			}

			resp.Msg = "Error determining state"

			ctx.JSON(http.StatusInternalServerError, resp)

			return
		}

		resp.BanType = bannedPerson.BanType

		var reason string

		switch {
		case bannedPerson.Reason == domain.Custom && bannedPerson.ReasonText != "":
			reason = bannedPerson.ReasonText
		case bannedPerson.Reason == domain.Custom && bannedPerson.ReasonText == "":
			reason = "Banned"
		default:
			reason = bannedPerson.Reason.String()
		}

		conf := s.configUsecase.Config()

		resp.Msg = fmt.Sprintf("Banned\nReason: %s\nAppeal: %s\nRemaining: %s", reason, conf.ExtURL(bannedPerson.BanSteam),
			time.Until(bannedPerson.ValidUntil).Round(time.Minute).String())

		ctx.JSON(http.StatusOK, resp)

		//goland:noinspection GoSwitchMissingCasesForIotaConsts
		switch resp.BanType {
		case domain.NoComm:
			slog.Info("Player muted", slog.Int64("sid64", sid.Int64()))
		case domain.Banned:
			// TODO log to discord
			slog.Info("Player dropped", slog.String("drop_type", "steam"),
				slog.Int64("sid64", sid.Int64()))
		}
	}
}

func (s *srcdsHandler) checkASN(ctx *gin.Context, steamID steamid.SteamID, addr netip.Addr, responseCtx context.Context, resp *CheckResponse) bool {
	details, errDetails := s.networkUsecase.QueryNetwork(ctx, addr)
	if errDetails == nil && details.Asn.ASNum > 0 {
		var asnBan domain.BanASN
		if errASNBan := s.banASNUsecase.GetByASN(responseCtx, int64(details.Asn.ASNum), &asnBan); errASNBan != nil {
			if !errors.Is(errASNBan, domain.ErrNoResult) {
				slog.Error("Failed to fetch asn bannedPerson", log.ErrAttr(errASNBan))
			}
			// Fail open
			return false
		}

		resp.BanType = domain.Banned
		resp.Msg = asnBan.Reason.String()
		ctx.JSON(http.StatusOK, resp)
		slog.Info("Player dropped", slog.String("drop_type", "asn"),
			slog.Int64("sid64", steamID.Int64()))

		return true
	}

	return false
}

func (s *srcdsHandler) checkIPBan(ctx *gin.Context, steamID steamid.SteamID, addr netip.Addr, responseCtx context.Context, resp *CheckResponse) bool {
	// Check IP first
	banNet, errGetBanNet := s.banNetUsecase.GetByAddress(responseCtx, addr)
	if errGetBanNet != nil {
		ctx.JSON(http.StatusInternalServerError, CheckResponse{
			BanType: domain.Unknown,
			Msg:     "Error determining state",
		})
		slog.Error("Could not get bannedPerson net results", log.ErrAttr(errGetBanNet))

		return true
	}

	if len(banNet) > 0 {
		resp.BanType = domain.Banned
		resp.Msg = "Banned"

		ctx.JSON(http.StatusOK, resp)

		slog.Info("Player dropped", slog.String("drop_type", "cidr"),
			slog.Int64("sid64", steamID.Int64()))

		return true
	}

	return false
}

func (s *srcdsHandler) checkNetBlockBan(ctx *gin.Context, steamID steamid.SteamID, addr netip.Addr, resp *CheckResponse) bool {
	if source, cidrBanned := s.networkUsecase.IsMatch(addr); cidrBanned {
		resp.BanType = domain.Network
		resp.Msg = "Network Range Banned.\nIf you using a VPN try disabling it"

		ctx.JSON(http.StatusOK, resp)
		slog.Info("Player network blocked", slog.Int64("sid64", steamID.Int64()),
			slog.String("source", source), slog.String("ip", addr.String()))

		return true
	}

	return false
}

func (s *srcdsHandler) checkGroupBan(ctx *gin.Context, steamID steamid.SteamID, resp *CheckResponse) bool {
	if groupID, banned := s.banGroupUsecase.IsMember(steamID); banned {
		resp.BanType = domain.Banned
		resp.Msg = fmt.Sprintf("Group Banned (gid: %d)", groupID.Int64())

		ctx.JSON(http.StatusOK, resp)
		slog.Info("Player dropped", slog.String("drop_type", "group"),
			slog.Int64("sid64", steamID.Int64()))

		return true
	}

	return false
}

func (s *srcdsHandler) checkFriendBan(ctx *gin.Context, steamID steamid.SteamID, resp *CheckResponse) bool {
	if parentFriendID, banned := s.banUsecase.IsFriendBanned(steamID); banned {
		resp.BanType = domain.Banned

		resp.Msg = fmt.Sprintf("Banned (sid: %d)", parentFriendID.Int64())

		ctx.JSON(http.StatusOK, resp)
		slog.Info("Player dropped", slog.String("drop_type", "friend"),
			slog.Int64("sid64", steamID.Int64()))

		return true
	}

	return false
}
