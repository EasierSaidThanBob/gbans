package api

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"github.com/leighmacdonald/gbans/internal/errs"
	"github.com/leighmacdonald/gbans/internal/model"
	"github.com/leighmacdonald/gbans/pkg/util"
	"github.com/leighmacdonald/steamid/v3/steamid"
	"go.uber.org/zap"
)

const ctxKeyUserProfile = "user_profile"

var (
	errInvalidParameter = errors.New("invalid parameter format")
	errPermissionDenied = errors.New("permission denied")
	errBadRequest       = errors.New("invalid request")
	errInternal         = errors.New("internal server error")
	errParamKeyMissing  = errors.New("param key not found")
	errParamParse       = errors.New("failed to parse param value")
	errParamInvalid     = errors.New("param value invalid")
)

type apiError struct {
	Message string `json:"message"`
}

func responseErr(ctx *gin.Context, statusCode int, err error) {
	userErr := "API Error"
	if err != nil {
		userErr = err.Error()
	}

	ctx.JSON(statusCode, apiError{Message: userErr})
}

func bind(ctx *gin.Context, log *zap.Logger, target any) bool {
	if errBind := ctx.BindJSON(&target); errBind != nil {
		responseErr(ctx, http.StatusBadRequest, errBadRequest)
		log.Error("Failed to bind request", zap.Error(errBind))

		return false
	}

	return true
}

func getSID64Param(c *gin.Context, key string) (steamid.SID64, error) {
	i, errGetParam := getInt64Param(c, key)
	if errGetParam != nil {
		return "", errGetParam
	}

	sid := steamid.New(i)
	if !sid.Valid() {
		return "", errs.ErrInvalidSID
	}

	return sid, nil
}

func getInt64Param(ctx *gin.Context, key string) (int64, error) {
	valueStr := ctx.Param(key)
	if valueStr == "" {
		return 0, fmt.Errorf("%w: %s", errParamKeyMissing, key)
	}

	value, valueErr := strconv.ParseInt(valueStr, 10, 64)
	if valueErr != nil {
		return 0, errParamParse
	}

	if value <= 0 {
		return 0, fmt.Errorf("%w: %s", errParamInvalid, key)
	}

	return value, nil
}

func getIntParam(ctx *gin.Context, key string) (int, error) {
	valueStr := ctx.Param(key)
	if valueStr == "" {
		return 0, fmt.Errorf("%w: %s", errParamKeyMissing, key)
	}

	return util.StringToInt(valueStr), nil
}

func getUUIDParam(ctx *gin.Context, key string) (uuid.UUID, error) {
	valueStr := ctx.Param(key)
	if valueStr == "" {
		return uuid.UUID{}, fmt.Errorf("%w: %s", errParamKeyMissing, key)
	}

	parsedUUID, errString := uuid.FromString(valueStr)
	if errString != nil {
		return uuid.UUID{}, errors.Join(errString, errParamParse)
	}

	return parsedUUID, nil
}

func serverFromCtx(ctx *gin.Context) int {
	serverIDUntyped, ok := ctx.Get("server_id")
	if !ok {
		return 0
	}

	serverID, castOk := serverIDUntyped.(int)
	if !castOk {
		return 0
	}

	return serverID
}

func contestFromCtx(ctx *gin.Context, env Env) (model.Contest, bool) {
	contestID, idErr := getUUIDParam(ctx, "contest_id")
	if idErr != nil {
		responseErr(ctx, http.StatusBadRequest, errBadRequest)

		return model.Contest{}, false
	}

	var contest model.Contest
	if errContests := env.Store().ContestByID(ctx, contestID, &contest); errContests != nil {
		responseErr(ctx, http.StatusInternalServerError, errInternal)

		return model.Contest{}, false
	}

	if !contest.Public && currentUserProfile(ctx).PermissionLevel < model.PModerator {
		responseErr(ctx, http.StatusForbidden, errs.ErrNotFound)

		return model.Contest{}, false
	}

	return contest, true
}

func newLazyResult(count int64, data any) LazyResult {
	if count == 0 {
		return LazyResult{0, []interface{}{}}
	}

	return LazyResult{Count: count, Data: data}
}

func New(ctx context.Context, env Env) *http.Server {
	conf := env.Config()

	httpServer := &http.Server{
		Addr:           conf.HTTP.Addr(),
		Handler:        createRouter(ctx, env),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if conf.HTTP.TLS {
		tlsVar := &tls.Config{
			// Only use curves which have assembly implementations
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519, // Go 1.8 only
			},
			MinVersion: tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			},
		}
		httpServer.TLSConfig = tlsVar
	}

	return httpServer
}

func currentUserProfile(ctx *gin.Context) model.UserProfile {
	maybePerson, found := ctx.Get(ctxKeyUserProfile)
	if !found {
		return model.NewUserProfile("")
	}

	person, ok := maybePerson.(model.UserProfile)
	if !ok {
		return model.NewUserProfile("")
	}

	return person
}

// checkPrivilege first checks if the steamId matches one of the provided allowedSteamIds, otherwise it will check
// if the user has appropriate privilege levels.
// Error responses are handled by this function, no further action needs to take place in the handlers.
func checkPrivilege(ctx *gin.Context, person model.UserProfile, allowedSteamIds steamid.Collection, minPrivilege model.Privilege) bool {
	for _, steamID := range allowedSteamIds {
		if steamID == person.SteamID {
			return true
		}
	}

	if person.PermissionLevel >= minPrivilege {
		return true
	}

	ctx.JSON(http.StatusForbidden, errPermissionDenied.Error())

	return false
}

type ResultsCount struct {
	Count int64 `json:"count"`
}

type LazyResult struct {
	Count int64 `json:"count"`
	Data  any   `json:"data"`
}
