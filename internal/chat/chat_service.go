package chat

import (
	"errors"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/leighmacdonald/gbans/internal/domain"
	"github.com/leighmacdonald/gbans/internal/http_helper"
	"go.uber.org/zap"
)

type ChatHandler struct {
	cu  domain.ChatUsecase
	log *zap.Logger
}

func NewChatHandler(log *zap.Logger, engine *gin.Engine, cu domain.ChatUsecase, ath domain.AuthUsecase) {
	handler := ChatHandler{
		cu:  cu,
		log: log.Named("chat"),
	}

	// authed
	authedGrp := engine.Group("/")
	{
		authed := authedGrp.Use(ath.AuthMiddleware(domain.PUser))
		authed.POST("/api/messages", handler.onAPIQueryMessages())
	}

	// mod
	modGrp := engine.Group("/")
	{
		mod := modGrp.Use(ath.AuthMiddleware(domain.PModerator))
		mod.GET("/api/message/:person_message_id/context/:padding", handler.onAPIQueryMessageContext())
	}
}

func (h ChatHandler) onAPIQueryMessages() gin.HandlerFunc {
	log := h.log.Named(runtime.FuncForPC(make([]uintptr, 10)[0]).Name())

	return func(ctx *gin.Context) {
		var req domain.ChatHistoryQueryFilter
		if !http_helper.Bind(ctx, log, &req) {
			return
		}

		if req.Limit <= 0 || req.Limit > 1000 {
			req.Limit = 50
		}

		user := http_helper.CurrentUserProfile(ctx)

		if user.PermissionLevel <= domain.PUser {
			req.Unrestricted = false
			beforeLimit := time.Now().Add(-time.Minute * 20)

			if req.DateEnd != nil && req.DateEnd.After(beforeLimit) {
				req.DateEnd = &beforeLimit
			}

			if req.DateEnd == nil {
				req.DateEnd = &beforeLimit
			}
		} else {
			req.Unrestricted = true
		}

		messages, count, errChat := h.cu.QueryChatHistory(ctx, req)
		if errChat != nil && !errors.Is(errChat, domain.ErrNoResult) {
			log.Error("Failed to query messages history",
				zap.Error(errChat), zap.String("sid", string(req.SourceID)))
			http_helper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		ctx.JSON(http.StatusOK, domain.NewLazyResult(count, messages))
	}
}

func (h ChatHandler) onAPIQueryMessageContext() gin.HandlerFunc {
	log := h.log.Named(runtime.FuncForPC(make([]uintptr, 10)[0]).Name())

	return func(ctx *gin.Context) {
		messageID, errMessageID := http_helper.GetInt64Param(ctx, "person_message_id")
		if errMessageID != nil {
			http_helper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrInvalidParameter)
			log.Debug("Got invalid person_message_id", zap.Error(errMessageID))

			return
		}

		padding, errPadding := http_helper.GetIntParam(ctx, "padding")
		if errPadding != nil {
			http_helper.ResponseErr(ctx, http.StatusBadRequest, domain.ErrBadRequest)
			log.Debug("Got invalid padding", zap.Error(errPadding))

			return
		}

		var msg domain.QueryChatHistoryResult
		if errMsg := h.cu.GetPersonMessage(ctx, messageID, &msg); errMsg != nil {
			http_helper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		messages, errQuery := h.cu.GetPersonMessageContext(ctx, msg.ServerID, messageID, padding)
		if errQuery != nil {
			http_helper.ResponseErr(ctx, http.StatusInternalServerError, domain.ErrInternal)

			return
		}

		ctx.JSON(http.StatusOK, messages)
	}
}