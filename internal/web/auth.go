package web

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/leighmacdonald/gbans/internal/app"
	"github.com/leighmacdonald/gbans/internal/config"
	"github.com/leighmacdonald/gbans/internal/consts"
	"github.com/leighmacdonald/gbans/internal/model"
	"github.com/leighmacdonald/gbans/internal/store"
	"github.com/leighmacdonald/gbans/pkg/util"
	"github.com/leighmacdonald/steamid/v2/steamid"
	"github.com/pkg/errors"
	"github.com/yohcop/openid-go"
	"go.uber.org/zap"
)

// noOpDiscoveryCache implements the DiscoveryCache interface and doesn't cache anything.
type noOpDiscoveryCache struct{}

// Put is a no op.
func (n *noOpDiscoveryCache) Put(_ string, _ openid.DiscoveredInfo) {}

// Get always returns nil.
func (n *noOpDiscoveryCache) Get(_ string) openid.DiscoveredInfo {
	return nil
}

var (
	nonceStore     = openid.NewSimpleNonceStore()
	discoveryCache = &noOpDiscoveryCache{}
)

func authServerMiddleWare(cookieKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims := &serverAuthClaims{}
		parsedToken, errParseClaims := jwt.ParseWithClaims(authHeader, claims, makeGetTokenKey(cookieKey))
		if errParseClaims != nil {
			if errors.Is(errParseClaims, jwt.ErrSignatureInvalid) {
				logger.Error("jwt signature invalid!", zap.Error(errParseClaims))
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if !parsedToken.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			logger.Error("Invalid jwt token parsed")
			return
		}
		if claims.ServerID <= 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			logger.Error("Invalid jwt claim server")
			return
		}
		var server store.Server
		if errGetServer := store.GetServer(ctx, claims.ServerID, &server); errGetServer != nil {
			logger.Error("Failed to load server during auth", zap.Error(errGetServer))
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		ctx.Next()
	}
}

func onGetLogout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO Logout key / mark as invalid manually
		logger.Error("onGetLogout Unimplemented")
		ctx.Redirect(http.StatusTemporaryRedirect, "/")
	}
}

func referral(ctx *gin.Context) string {
	referralURL, found := ctx.GetQuery("return_url")
	if !found {
		referralURL = "/"
	}
	return referralURL
}

func onOAuthDiscordCallback(conf *config.Config) gin.HandlerFunc {
	client := util.NewHTTPClient()
	type accessTokenResp struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
		TokenType    string `json:"token_type"`
	}
	type discordUserDetail struct {
		ID               string      `json:"id"`
		Username         string      `json:"username"`
		Avatar           string      `json:"avatar"`
		AvatarDecoration interface{} `json:"avatar_decoration"`
		Discriminator    string      `json:"discriminator"`
		PublicFlags      int         `json:"public_flags"`
		Flags            int         `json:"flags"`
		Banner           interface{} `json:"banner"`
		BannerColor      interface{} `json:"banner_color"`
		AccentColor      interface{} `json:"accent_color"`
		Locale           string      `json:"locale"`
		MfaEnabled       bool        `json:"mfa_enabled"`
		PremiumType      int         `json:"premium_type"`
	}

	fetchDiscordID := func(ctx context.Context, accessToken string) (string, error) {
		req, errReq := http.NewRequestWithContext(ctx, http.MethodGet, "https://discord.com/api/users/@me", nil)
		if errReq != nil {
			return "", errReq
		}
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		resp, errResp := client.Do(req)
		if errResp != nil {
			return "", errResp
		}
		defer func() {
			_ = resp.Body.Close()
		}()
		b, errBody := io.ReadAll(resp.Body)
		if errBody != nil {
			return "", errBody
		}
		var details discordUserDetail
		if errJSON := json.Unmarshal(b, &details); errJSON != nil {
			return "", errJSON
		}
		return details.ID, nil
	}

	fetchToken := func(ctx context.Context, code string) (string, error) {
		form := url.Values{}
		form.Set("client_id", conf.Discord.AppID)
		form.Set("client_secret", conf.Discord.AppSecret)
		form.Set("redirect_uri", conf.ExtURL("/login/discord"))
		form.Set("code", code)
		form.Set("grant_type", "authorization_code")
		form.Set("scope", "identify")
		req, errReq := http.NewRequestWithContext(ctx, http.MethodPost, "https://discord.com/api/oauth2/token", strings.NewReader(form.Encode()))
		if errReq != nil {
			return "", errReq
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		resp, errResp := client.Do(req)
		if errResp != nil {
			return "", errResp
		}
		defer func() {
			_ = resp.Body.Close()
		}()
		body, errBody := io.ReadAll(resp.Body)
		if errBody != nil {
			return "", errBody
		}
		var atr accessTokenResp
		if errJSON := json.Unmarshal(body, &atr); errJSON != nil {
			return "", errJSON
		}
		return atr.AccessToken, nil
	}

	return func(ctx *gin.Context) {
		code := ctx.Query("code")
		if code == "" {
			responseErr(ctx, http.StatusBadRequest, nil)
			return
		}
		token, errToken := fetchToken(ctx, code)
		if errToken != nil {
			responseErr(ctx, http.StatusBadRequest, nil)
			return
		}
		discordID, errID := fetchDiscordID(ctx, token)
		if errID != nil {
			responseErr(ctx, http.StatusBadRequest, nil)
			return
		}
		if discordID == "" {
			responseErr(ctx, http.StatusInternalServerError, nil)
			return
		}
		var dp store.Person
		if errDp := store.GetPersonByDiscordID(ctx, discordID, &dp); errDp != nil {
			if !errors.Is(errDp, store.ErrNoResult) {
				responseErr(ctx, http.StatusInternalServerError, nil)
				return
			}
		}
		if dp.DiscordID != "" {
			responseErr(ctx, http.StatusConflict, nil)
			return
		}
		sid := currentUserProfile(ctx).SteamID
		var sp store.Person
		if errPerson := app.PersonBySID(ctx, sid, &sp); errPerson != nil {
			responseErr(ctx, http.StatusInternalServerError, nil)
			return
		}
		sp.DiscordID = discordID
		if errSave := store.SavePerson(ctx, &sp); errSave != nil {
			responseErr(ctx, http.StatusInternalServerError, nil)
			return
		}
		responseOK(ctx, http.StatusInternalServerError, nil)
		logger.Info("Discord account linked successfully",
			zap.String("discord_id", discordID), zap.Int64("sid64", sid.Int64()))
	}
}

func onOpenIDCallback(conf *config.Config) gin.HandlerFunc {
	oidRx := regexp.MustCompile(`^https://steamcommunity\.com/openid/id/(\d+)$`)
	return func(ctx *gin.Context) {
		var idStr string
		referralURL := referral(ctx)
		fullURL := conf.General.ExternalURL + ctx.Request.URL.String()
		if conf.Debug.SkipOpenIDValidation {
			// Pull the sid out of the query without doing a signature check
			values, errParse := url.Parse(fullURL)
			if errParse != nil {
				logger.Error("Failed to parse url", zap.Error(errParse))
				ctx.Redirect(302, referralURL)
				return
			}
			idStr = values.Query().Get("openid.identity")
		} else {
			id, errVerify := openid.Verify(fullURL, discoveryCache, nonceStore)
			if errVerify != nil {
				logger.Error("Error verifying openid auth response", zap.Error(errVerify))
				ctx.Redirect(302, referralURL)
				return
			}
			idStr = id
		}
		match := oidRx.FindStringSubmatch(idStr)
		if match == nil || len(match) != 2 {
			ctx.Redirect(302, referralURL)
			return
		}
		sid, errDecodeSid := steamid.SID64FromString(match[1])
		if errDecodeSid != nil {
			logger.Error("Received invalid steamid", zap.Error(errDecodeSid))
			ctx.Redirect(302, referralURL)
			return
		}
		person := store.NewPerson(sid)
		if errGetProfile := app.PersonBySID(ctx, sid, &person); errGetProfile != nil {
			logger.Error("Failed to fetch user profile", zap.Error(errGetProfile))
			ctx.Redirect(302, referralURL)
			return
		}
		accessToken, refreshToken, errToken := makeTokens(ctx, conf.HTTP.CookieKey, sid)
		if errToken != nil {
			ctx.Redirect(302, referralURL)
			logger.Error("Failed to create access token pair", zap.Error(errToken))
			return
		}
		parsedURL, errParse := url.Parse("/login/success")
		if errParse != nil {
			ctx.Redirect(302, referralURL)
			return
		}
		query := parsedURL.Query()
		query.Set("refresh", refreshToken)
		query.Set("token", accessToken)
		query.Set("next_url", referralURL)
		parsedURL.RawQuery = query.Encode()
		ctx.Redirect(302, parsedURL.String())
		logger.Info("User logged in",
			zap.Int64("sid64", sid.Int64()),
			zap.String("name", person.PersonaName),
			zap.Int("permission_level", int(person.PermissionLevel)))
	}
}

func makeTokens(ctx *gin.Context, cookieKey string, sid steamid.SID64) (string, string, error) {
	accessToken, errJWT := newUserJWT(sid, cookieKey)
	if errJWT != nil {
		return "", "", errors.Wrap(errJWT, "Failed to create new access token")
	}
	ipAddr := net.ParseIP(ctx.ClientIP())
	refreshToken := store.NewPersonAuth(sid, ipAddr)
	if errAuth := store.GetPersonAuth(ctx, sid, ipAddr, &refreshToken); errAuth != nil {
		if !errors.Is(errAuth, store.ErrNoResult) {
			return "", "", errors.Wrap(errAuth, "Failed to fetch refresh token")
		}
		if createErr := store.SavePersonAuth(ctx, &refreshToken); createErr != nil {
			return "", "", errors.Wrap(errAuth, "Failed to create new refresh token")
		}
	}
	return accessToken, refreshToken.RefreshToken, nil
}

func makeGetTokenKey(cookieKey string) func(_ *jwt.Token) (any, error) {
	return func(_ *jwt.Token) (any, error) {
		return []byte(cookieKey), nil
	}
}

// onTokenRefresh handles generating new token pairs to access the api
// NOTE: All error code paths must return 401 (Unauthorized).
func onTokenRefresh(conf *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var rt userToken
		if errBind := ctx.BindJSON(&rt); errBind != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			logger.Error("Malformed user token", zap.Error(errBind))
			return
		}
		if rt.RefreshToken == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		var auth store.PersonAuth
		if authError := store.GetPersonAuthByRefreshToken(ctx, rt.RefreshToken, &auth); authError != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		newAccessToken, newRefreshToken, errToken := makeTokens(ctx, conf.HTTP.CookieKey, auth.SteamID)
		if errToken != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			logger.Error("Failed to create access token pair", zap.Error(errToken))
			return
		}
		responseOK(ctx, http.StatusOK, userToken{
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken,
		})
	}
}

type userToken struct {
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken"`
}

type personAuthClaims struct {
	SteamID int64 `json:"steam_id"`
	jwt.StandardClaims
}

type serverAuthClaims struct {
	ServerID int `json:"server_id"`
	jwt.StandardClaims
}

const authTokenLifetimeDuration = time.Hour * 24 * 30 // 1 month

func newUserJWT(steamID steamid.SID64, cookieKey string) (string, error) {
	t0 := config.Now()
	claims := &personAuthClaims{
		SteamID: steamID.Int64(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: t0.Add(authTokenLifetimeDuration).Unix(),
			IssuedAt:  t0.Unix(),
		},
	}
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, errSigned := tokenWithClaims.SignedString([]byte(cookieKey))
	if errSigned != nil {
		return "", errors.Wrap(errSigned, "Failed create signed string")
	}
	return signedToken, nil
}

func newServerJWT(serverID int, cookieKey string) (string, error) {
	t0 := config.Now()
	claims := &serverAuthClaims{
		ServerID: serverID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: t0.Add(authTokenLifetimeDuration).Unix(),
			IssuedAt:  t0.Unix(),
		},
	}
	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, errSigned := tokenWithClaims.SignedString([]byte(cookieKey))
	if errSigned != nil {
		return "", errors.Wrap(errSigned, "Failed create signed string")
	}
	return signedToken, nil
}

// authMiddleware handles client authentication to the HTTP & websocket api.
// websocket clients must pass the key as a query parameter called "token".
func authMiddleware(conf *config.Config, level consts.Privilege) gin.HandlerFunc {
	type header struct {
		Authorization string `header:"Authorization"`
	}
	return func(ctx *gin.Context) {
		var token string
		if ctx.FullPath() == "/ws" {
			token = ctx.Query("token")
		} else {
			hdr := header{}
			if errBind := ctx.ShouldBindHeader(&hdr); errBind != nil {
				ctx.AbortWithStatus(http.StatusForbidden)
				return
			}
			pcs := strings.Split(hdr.Authorization, " ")
			if len(pcs) != 2 && level >= consts.PUser {
				ctx.AbortWithStatus(http.StatusForbidden)
				return
			}
			token = pcs[1]
		}
		if level >= consts.PUser {
			sid, errFromToken := sid64FromJWTToken(token, conf.HTTP.CookieKey)
			if errFromToken != nil {
				if errors.Is(errFromToken, consts.ErrExpired) {
					ctx.AbortWithStatus(http.StatusUnauthorized)
					return
				}
				logger.Error("Failed to load sid from access token", zap.Error(errFromToken))
				ctx.AbortWithStatus(http.StatusForbidden)
				return
			}
			loggedInPerson := store.NewPerson(sid)
			if errGetPerson := app.PersonBySID(ctx, sid, &loggedInPerson); errGetPerson != nil {
				logger.Error("Failed to load person during auth", zap.Error(errGetPerson))
				ctx.AbortWithStatus(http.StatusForbidden)
				return
			}
			if level > loggedInPerson.PermissionLevel {
				ctx.AbortWithStatus(http.StatusForbidden)
				return
			}
			bp := store.NewBannedPerson()
			if errBan := store.GetBanBySteamID(ctx, sid, &bp, false); errBan != nil {
				if !errors.Is(errBan, store.ErrNoResult) {
					logger.Error("Failed to fetch authed user ban", zap.Error(errBan))
				}
			}
			profile := model.UserProfile{
				SteamID:         loggedInPerson.SteamID,
				CreatedOn:       loggedInPerson.CreatedOn,
				UpdatedOn:       loggedInPerson.UpdatedOn,
				PermissionLevel: loggedInPerson.PermissionLevel,
				DiscordID:       loggedInPerson.DiscordID,
				Name:            loggedInPerson.PersonaName,
				Avatar:          loggedInPerson.Avatar,
				AvatarFull:      loggedInPerson.AvatarFull,
				Muted:           loggedInPerson.Muted,
				BanID:           bp.Ban.BanID,
			}
			ctx.Set(ctxKeyUserProfile, profile)
		}
		ctx.Next()
	}
}

func sid64FromJWTToken(token string, cookieKey string) (steamid.SID64, error) {
	claims := &personAuthClaims{}
	tkn, errParseClaims := jwt.ParseWithClaims(token, claims, makeGetTokenKey(cookieKey))
	if errParseClaims != nil {
		if errors.Is(errParseClaims, jwt.ErrSignatureInvalid) {
			return 0, consts.ErrAuthentication
		}
		var e *jwt.ValidationError
		ok := errors.Is(errParseClaims, e)
		if ok && e.Errors == jwt.ValidationErrorExpired {
			return 0, consts.ErrExpired
		}
		return 0, consts.ErrAuthentication
	}
	if !tkn.Valid {
		return 0, consts.ErrAuthentication
	}
	sid := steamid.SID64(claims.SteamID)
	if !sid.Valid() {
		return 0, consts.ErrAuthentication
	}
	return sid, nil
}