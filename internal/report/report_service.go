package report

import (
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leighmacdonald/gbans/internal/discord"
	"github.com/leighmacdonald/gbans/internal/domain"
	"github.com/leighmacdonald/gbans/internal/httphelper"
	"github.com/leighmacdonald/gbans/internal/thirdparty"
	"github.com/leighmacdonald/gbans/pkg/fp"
	"github.com/leighmacdonald/steamid/v3/steamid"
	"go.uber.org/zap"
)

type ReportHandler struct {
	log            *zap.Logger
	reportUsecase  domain.ReportUsecase
	configUsecase  domain.ConfigUsecase
	discordUsecase domain.DiscordUsecase
	personUsecase  domain.PersonUsecase
}

func NewReportHandler(log *zap.Logger, engine *gin.Engine, reportUsecase domain.ReportUsecase, configUsecase domain.ConfigUsecase,
	discordUsecase domain.DiscordUsecase, personUsecase domain.PersonUsecase, authUsecase domain.AuthUsecase,
) {
	handler := ReportHandler{
		log:            log.Named("report"),
		reportUsecase:  reportUsecase,
		configUsecase:  configUsecase,
		discordUsecase: discordUsecase,
		personUsecase:  personUsecase,
	}

	// auth
	authedGrp := engine.Group("/")
	{
		authed := authedGrp.Use(authUsecase.AuthMiddleware(domain.PUser))
		authed.POST("/api/report", handler.onAPIPostReportCreate())
		authed.GET("/api/report/:report_id", handler.onAPIGetReport())
		authed.POST("/api/reports", handler.onAPIGetReports())
		authed.POST("/api/report_status/:report_id", handler.onAPISetReportStatus())
		authed.GET("/api/report/:report_id/messages", handler.onAPIGetReportMessages())
		authed.POST("/api/report/:report_id/messages", handler.onAPIPostReportMessage())
		authed.POST("/api/report/message/:report_message_id", handler.onAPIEditReportMessage())
		authed.DELETE("/api/report/message/:report_message_id", handler.onAPIDeleteReportMessage())
	}
	// mod
	modGrp := engine.Group("/")
	{
		mod := modGrp.Use(authUsecase.AuthMiddleware(domain.PModerator))
		mod.POST("/api/report/:report_id/state", handler.onAPIPostBanState())
	}
}

func (h ReportHandler) onAPIPostBanState() gin.HandlerFunc {
	// TODO doesnt do anything
	return func(ctx *gin.Context) {
		reportID, errID := httphelper.GetInt64Param(ctx, "report_id")
		if errID != nil || reportID <= 0 {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInvalidParameter)

			return
		}

		report, errReport := h.reportUsecase.GetReport(ctx, httphelper.CurrentUserProfile(ctx), reportID)
		if errReport != nil {
			httphelper.ErrorHandled(ctx, errReport)

			return
		}

		ctx.JSON(http.StatusOK, report)

		h.discordUsecase.SendPayload(domain.ChannelModLog, discord.EditBanAppealStatusMessage())
	}
}

func (h ReportHandler) onAPIPostReportCreate() gin.HandlerFunc {
	log := h.log.Named(runtime.FuncForPC(make([]uintptr, 10)[0]).Name())

	return func(ctx *gin.Context) {
		currentUser := httphelper.CurrentUserProfile(ctx)

		var req domain.CreateReportReq
		if !httphelper.Bind(ctx, log, &req) {
			return
		}

		if req.Description == "" || len(req.Description) < 10 {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, fmt.Errorf("%w: description", domain.ErrParamInvalid))

			return
		}

		// ServerStore initiated requests will have a sourceID set by the server
		// Web based reports the source should not be set, the reporter will be taken from the
		// current session information instead
		if req.SourceID == "" {
			req.SourceID = domain.StringSID(currentUser.SteamID.String())
		}

		sourceID, errSourceID := req.SourceID.SID64(ctx)
		if errSourceID != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrSourceID)
			log.Error("Invalid steam_id", zap.Error(errSourceID))

			return
		}

		targetID, errTargetID := req.TargetID.SID64(ctx)
		if errTargetID != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrTargetID)
			log.Error("Invalid target_id", zap.Error(errTargetID))

			return
		}

		if sourceID == targetID {
			httphelper.ResponseErr(ctx, http.StatusConflict, domain.ErrSelfReport)

			return
		}

		personSource, errSource := h.personUsecase.GetPersonBySteamID(ctx, sourceID)
		if errSource != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)
			log.Error("Could not load player profile", zap.Error(errSource))

			return
		}

		personTarget, errTarget := h.personUsecase.GetOrCreatePersonBySteamID(ctx, targetID)
		if errTarget != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)
			log.Error("Could not load player profile", zap.Error(errTarget))

			return
		}

		if personTarget.Expired() {
			if err := thirdparty.UpdatePlayerSummary(ctx, &personTarget); err != nil {
				log.Error("Failed to update target player", zap.Error(err))
			} else {
				if errSave := h.personUsecase.SavePerson(ctx, &personTarget); errSave != nil {
					log.Error("Failed to save target player update", zap.Error(err))
				}
			}
		}

		// Ensure the user doesn't already have an open report against the user
		existing, errReports := h.reportUsecase.GetReportBySteamID(ctx, personSource.SteamID, targetID)
		if errReports != nil {
			if !errors.Is(errReports, domain.ErrNoResult) {
				log.Error("Failed to query reports by steam id", zap.Error(errReports))
				httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

				return
			}
		}

		if existing.ReportID > 0 {
			httphelper.ResponseErr(ctx, http.StatusConflict, domain.ErrReportExists)

			return
		}

		// TODO encapsulate all operations in single tx
		report := domain.NewReport()
		report.SourceID = sourceID
		report.ReportStatus = domain.Opened
		report.Description = req.Description
		report.TargetID = targetID
		report.Reason = req.Reason
		report.ReasonText = req.ReasonText
		parts := strings.Split(req.DemoName, "/")
		report.DemoName = parts[len(parts)-1]
		report.DemoTick = req.DemoTick
		report.PersonMessageID = req.PersonMessageID

		if errReportSave := h.reportUsecase.SaveReport(ctx, &report); errReportSave != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)
			log.Error("Failed to save report", zap.Error(errReportSave))

			return
		}

		ctx.JSON(http.StatusCreated, report)

		log.Info("New report created successfully", zap.Int64("report_id", report.ReportID))

		conf := h.configUsecase.Config()

		if !conf.Discord.Enabled {
			return
		}

		demoURL := ""

		if report.DemoName != "" {
			demoURL = conf.ExtURLRaw("/demos/name/%s", report.DemoName)
		}

		msg := discord.NewInGameReportResponse(report, conf.ExtURL(report), currentUser, conf.ExtURL(currentUser), demoURL)

		h.discordUsecase.SendPayload(domain.ChannelModLog, msg)
	}
}

func (h ReportHandler) onAPIGetReport() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reportID, errParam := httphelper.GetInt64Param(ctx, "report_id")
		if errParam != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInvalidParameter)

			return
		}

		report, errReport := h.reportUsecase.GetReport(ctx, httphelper.CurrentUserProfile(ctx), reportID)
		if errReport != nil {
			httphelper.ErrorHandled(ctx, errReport)

			return
		}

		ctx.JSON(http.StatusOK, report)
	}
}

type reportWithAuthor struct {
	Author  domain.Person `json:"author"`
	Subject domain.Person `json:"subject"`
	domain.Report
}

func (h ReportHandler) onAPIGetReports() gin.HandlerFunc {
	log := h.log.Named(runtime.FuncForPC(make([]uintptr, 10)[0]).Name())

	return func(ctx *gin.Context) {
		user := httphelper.CurrentUserProfile(ctx)

		var req domain.ReportQueryFilter
		if !httphelper.Bind(ctx, log, &req) {
			return
		}

		if req.Limit <= 0 && req.Limit > 100 {
			req.Limit = 25
		}

		// Make sure the person requesting is either a moderator, or a user
		// only able to request their own reports
		var sourceID steamid.SID64

		if user.PermissionLevel < domain.PModerator {
			sourceID = user.SteamID
		} else if req.SourceID != "" {
			sid, errSourceID := req.SourceID.SID64(ctx)
			if errSourceID != nil {
				httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrBadRequest)

				return
			}

			sourceID = sid
		}

		if sourceID.Valid() {
			req.SourceID = domain.StringSID(sourceID.String())
		}

		var userReports []reportWithAuthor

		reports, count, errReports := h.reportUsecase.GetReports(ctx, req)
		if errReports != nil {
			if errors.Is(errReports, domain.ErrNoResult) {
				ctx.JSON(http.StatusNoContent, nil)

				return
			}

			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		var authorIds steamid.Collection
		for _, report := range reports {
			authorIds = append(authorIds, report.SourceID)
		}

		authors, errAuthors := h.personUsecase.GetPeopleBySteamID(ctx, fp.Uniq(authorIds))
		if errAuthors != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		authorMap := authors.AsMap()

		var subjectIds steamid.Collection
		for _, report := range reports {
			subjectIds = append(subjectIds, report.TargetID)
		}

		subjects, errSubjects := h.personUsecase.GetPeopleBySteamID(ctx, fp.Uniq(subjectIds))
		if errSubjects != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		subjectMap := subjects.AsMap()

		for _, report := range reports {
			userReports = append(userReports, reportWithAuthor{
				Author:  authorMap[report.SourceID],
				Report:  report,
				Subject: subjectMap[report.TargetID],
			})
		}

		if userReports == nil {
			userReports = []reportWithAuthor{}
		}

		ctx.JSON(http.StatusOK, domain.NewLazyResult(count, userReports))
	}
}

func (h ReportHandler) onAPISetReportStatus() gin.HandlerFunc {
	type stateUpdateReq struct {
		Status domain.ReportStatus `json:"status"`
	}

	log := h.log.Named(runtime.FuncForPC(make([]uintptr, 10)[0]).Name())

	return func(ctx *gin.Context) {
		reportID, errParam := httphelper.GetInt64Param(ctx, "report_id")
		if errParam != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInvalidParameter)

			return
		}

		var req stateUpdateReq
		if !httphelper.Bind(ctx, log, &req) {
			return
		}

		report, errGet := h.reportUsecase.GetReport(ctx, httphelper.CurrentUserProfile(ctx), reportID)
		if errGet != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)
			log.Error("Failed to get report to set state", zap.Error(errGet))

			return
		}

		if report.ReportStatus == req.Status {
			ctx.JSON(http.StatusConflict, domain.ErrDuplicate)

			return
		}

		original := report.ReportStatus

		report.ReportStatus = req.Status
		if errSave := h.reportUsecase.SaveReport(ctx, &report); errSave != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)
			log.Error("Failed to save report state", zap.Error(errSave))

			return
		}

		ctx.JSON(http.StatusAccepted, nil)
		log.Info("Report status changed",
			zap.Int64("report_id", report.ReportID),
			zap.String("from_status", original.String()),
			zap.String("to_status", report.ReportStatus.String()))
		// discord.SendDiscord(model.NotificationPayload{
		//	Sids:     steamid.Collection{report.SourceID},
		//	Severity: db.SeverityInfo,
		//	Message:  "Report status updated",
		//	Link:     report.ToURL(),
		// })
	} //nolint:wsl
}

func (h ReportHandler) onAPIGetReportMessages() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reportID, errParam := httphelper.GetInt64Param(ctx, "report_id")
		if errParam != nil {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInvalidParameter)

			return
		}

		report, errGetReport := h.reportUsecase.GetReport(ctx, httphelper.CurrentUserProfile(ctx), reportID)
		if errGetReport != nil {
			httphelper.ResponseErr(ctx, http.StatusNotFound, domain.ErrNotFound)

			return
		}

		if !httphelper.HasPrivilege(httphelper.CurrentUserProfile(ctx), steamid.Collection{report.SourceID, report.TargetID}, domain.PModerator) {
			return
		}

		reportMessages, errGetReportMessages := h.reportUsecase.GetReportMessages(ctx, reportID)
		if errGetReportMessages != nil {
			httphelper.ResponseErr(ctx, http.StatusNotFound, domain.ErrPlayerNotFound)

			return
		}

		if reportMessages == nil {
			reportMessages = []domain.ReportMessage{}
		}

		ctx.JSON(http.StatusOK, reportMessages)
	}
}

func (h ReportHandler) onAPIPostReportMessage() gin.HandlerFunc {
	type newMessage struct {
		Message string `json:"message"`
	}

	log := h.log.Named(runtime.FuncForPC(make([]uintptr, 10)[0]).Name())

	return func(ctx *gin.Context) {
		reportID, errID := httphelper.GetInt64Param(ctx, "report_id")
		if errID != nil || reportID == 0 {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInvalidParameter)

			return
		}

		var req newMessage
		if !httphelper.Bind(ctx, log, &req) {
			return
		}

		if req.Message == "" {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrBadRequest)

			return
		}

		report, errReport := h.reportUsecase.GetReport(ctx, httphelper.CurrentUserProfile(ctx), reportID)
		if errReport != nil {
			if errors.Is(errReport, domain.ErrNoResult) {
				httphelper.ResponseErr(ctx, http.StatusNotFound, domain.ErrNotFound)

				return
			}

			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)
			log.Error("Failed to load report", zap.Error(errReport))

			return
		}

		person := httphelper.CurrentUserProfile(ctx)
		msg := domain.NewReportMessage(reportID, person.SteamID, req.Message)

		if errSave := h.reportUsecase.SaveReportMessage(ctx, &msg); errSave != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)
			log.Error("Failed to save report message", zap.Error(errSave))

			return
		}

		report.UpdatedOn = time.Now()

		if errSave := h.reportUsecase.SaveReport(ctx, &report); errSave != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)
			log.Error("Failed to update report activity", zap.Error(errSave))

			return
		}

		ctx.JSON(http.StatusCreated, msg)

		conf := h.configUsecase.Config()

		h.discordUsecase.SendPayload(domain.ChannelModLog,
			discord.NewReportMessageResponse(msg.MessageMD, conf.ExtURL(report), person, conf.ExtURL(person)))
	}
}

func (h ReportHandler) onAPIEditReportMessage() gin.HandlerFunc {
	type editMessage struct {
		BodyMD string `json:"body_md"`
	}

	log := h.log.Named(runtime.FuncForPC(make([]uintptr, 10)[0]).Name())

	return func(ctx *gin.Context) {
		reportMessageID, errID := httphelper.GetInt64Param(ctx, "report_message_id")
		if errID != nil || reportMessageID == 0 {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInvalidParameter)

			return
		}

		existing, errExist := h.reportUsecase.GetReportMessageByID(ctx, reportMessageID)
		if errExist != nil {
			if errors.Is(errExist, domain.ErrNoResult) {
				httphelper.ResponseErr(ctx, http.StatusNotFound, domain.ErrPlayerNotFound)

				return
			}

			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		curUser := httphelper.CurrentUserProfile(ctx)
		if !httphelper.HasPrivilege(curUser, steamid.Collection{existing.AuthorID}, domain.PModerator) {
			return
		}

		var req editMessage
		if !httphelper.Bind(ctx, log, &req) {
			return
		}

		if req.BodyMD == "" {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrBadRequest)

			return
		}

		if req.BodyMD == existing.MessageMD {
			httphelper.ResponseErr(ctx, http.StatusConflict, domain.ErrDuplicate)

			return
		}

		existing.MessageMD = req.BodyMD
		if errSave := h.reportUsecase.SaveReportMessage(ctx, &existing); errSave != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)
			log.Error("Failed to save report message", zap.Error(errSave))

			return
		}

		ctx.JSON(http.StatusCreated, req)

		conf := h.configUsecase.Config()

		msg := discord.EditReportMessageResponse(req.BodyMD, existing.MessageMD,
			conf.ExtURLRaw("/report/%d", existing.ReportID), curUser, conf.ExtURL(curUser))

		h.discordUsecase.SendPayload(domain.ChannelModLog, msg)
	}
}

func (h ReportHandler) onAPIDeleteReportMessage() gin.HandlerFunc {
	log := h.log.Named(runtime.FuncForPC(make([]uintptr, 10)[0]).Name())

	return func(ctx *gin.Context) {
		reportMessageID, errID := httphelper.GetInt64Param(ctx, "report_message_id")
		if errID != nil || reportMessageID == 0 {
			httphelper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInvalidParameter)

			return
		}

		existing, errExist := h.reportUsecase.GetReportMessageByID(ctx, reportMessageID)
		if errExist != nil {
			if errors.Is(errExist, domain.ErrNoResult) {
				httphelper.ResponseErr(ctx, http.StatusNotFound, domain.ErrNotFound)

				return
			}

			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		curUser := httphelper.CurrentUserProfile(ctx)
		if !httphelper.HasPrivilege(curUser, steamid.Collection{existing.AuthorID}, domain.PModerator) {
			return
		}

		existing.Deleted = true
		if errSave := h.reportUsecase.SaveReportMessage(ctx, &existing); errSave != nil {
			httphelper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)
			log.Error("Failed to save report message", zap.Error(errSave))

			return
		}

		ctx.JSON(http.StatusNoContent, nil)

		conf := h.configUsecase.Config()

		h.discordUsecase.SendPayload(domain.ChannelModLog, discord.DeleteReportMessage(existing, curUser, conf.ExtURL(curUser)))
	}
}
