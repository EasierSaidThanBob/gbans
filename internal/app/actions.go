package app

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/leighmacdonald/gbans/internal/config"
	"github.com/leighmacdonald/gbans/internal/consts"
	"github.com/leighmacdonald/gbans/internal/query"
	"github.com/leighmacdonald/gbans/internal/state"
	"github.com/leighmacdonald/gbans/internal/store"
	"github.com/leighmacdonald/gbans/pkg/discordutil"
	"github.com/leighmacdonald/gbans/pkg/util"
	"github.com/leighmacdonald/steamid/v2/steamid"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"strings"
)

func (app *App) SendDiscordPayload(payload discordutil.Payload) {
	app.logger.Debug("Sending discordutil payload",
		zap.String("channel_id", payload.ChannelId),
		zap.Bool("enabled", config.Discord.PublicLogChannelEnable))
	if config.Discord.PublicLogChannelEnable {
		select {
		case app.discordSendMsg <- payload:
		default:
			app.logger.Error("Cannot send discordutil payload, discordSendMsg channel full")
		}
	}
}

// Kick will kick the steam id from whatever server it is connected to.
func (app *App) Kick(ctx context.Context, origin store.Origin, target store.StringSID, author store.StringSID,
	reason store.Reason, playerInfo *state.PlayerInfo) error {
	authorSid64, errAid := author.SID64()
	if errAid != nil {
		return errAid
	}
	// kick the user if they currently are playing on a server
	var foundPI state.PlayerInfo
	if errFind := app.Find(ctx, target, "", &foundPI); errFind != nil {
		return errFind
	}
	if foundPI.Valid && foundPI.InGame {
		_, errExecRCON := query.ExecRCON(ctx, *foundPI.Server, fmt.Sprintf("sm_kick #%d %s", foundPI.Player.UserID, reason))
		if errExecRCON != nil {
			app.logger.Error("Failed to kick user after ban", zap.Error(errExecRCON))
			return errExecRCON
		}
		app.logger.Info("User Kicked", zap.String("origin", origin.String()),
			zap.String("target", string(target)), zap.String("author", util.SanitizeLog(authorSid64.String())))
	}
	if playerInfo != nil {
		*playerInfo = foundPI
	}

	return nil
}

// Silence will gag & mute a player
func (app *App) Silence(ctx context.Context, origin store.Origin, target store.StringSID, author store.StringSID,
	reason store.Reason, playerInfo *state.PlayerInfo) error {
	_, errAid := author.SID64()
	if errAid != nil {
		return errAid
	}
	// kick the user if they currently are playing on a server
	var foundPI state.PlayerInfo
	if errFind := app.Find(ctx, target, "", &foundPI); errFind != nil {
		return errFind
	}
	if foundPI.Valid && foundPI.InGame {
		_, errExecRCON := query.ExecRCON(
			ctx,
			*foundPI.Server,
			fmt.Sprintf(`sm_silence "#%s" %s`, steamid.SID64ToSID(foundPI.Player.SID), reason),
		)
		if errExecRCON != nil {
			app.logger.Error("Failed to kick user after ban", zap.Error(errExecRCON))
			return errExecRCON
		}
	}
	if playerInfo != nil {
		*playerInfo = foundPI
	}

	return nil
}

// SetSteam is used to associate a discordutil user with either steam id. This is used
// instead of requiring users to link their steam account to discordutil itself. It also
// means the bot does not require more privileged intents.
func (app *App) SetSteam(ctx context.Context, sid64 steamid.SID64, discordId string) error {
	newPerson := store.NewPerson(sid64)
	if errGetPerson := app.store.GetOrCreatePersonBySteamID(ctx, sid64, &newPerson); errGetPerson != nil || !sid64.Valid() {
		return consts.ErrInvalidSID
	}
	if (newPerson.DiscordID) != "" {
		return errors.Errorf("Discord account already linked to steam account: %d", newPerson.SteamID.Int64())
	}
	newPerson.DiscordID = discordId
	if errSavePerson := app.store.SavePerson(ctx, &newPerson); errSavePerson != nil {
		return consts.ErrInternal
	}
	app.logger.Info("Discord steamid set", zap.Int64("sid64", sid64.Int64()), zap.String("discordId", discordId))
	return nil
}

// Say is used to send a message to the server via sm_say
func (app *App) Say(ctx context.Context, author steamid.SID64, serverName string, message string) error {
	var server store.Server
	if errGetServer := app.store.GetServerByName(ctx, serverName, &server); errGetServer != nil {
		return errors.Errorf("Failed to fetch server: %s", serverName)
	}
	msg := fmt.Sprintf(`sm_say %s`, message)
	rconResponse, errExecRCON := query.ExecRCON(ctx, server, msg)
	if errExecRCON != nil {
		return errExecRCON
	}
	responsePieces := strings.Split(rconResponse, "\n")
	if len(responsePieces) < 2 {
		return errors.Errorf("Invalid response")
	}
	app.logger.Info("Server message sent", zap.Int64("author", author.Int64()), zap.String("msg", message))
	return nil
}

// CSay is used to send a centered message to the server via sm_csay
func (app *App) CSay(ctx context.Context, author steamid.SID64, serverName string, message string) error {
	var (
		servers []store.Server
		err     error
	)
	if serverName == "*" {
		servers, err = app.store.GetServers(ctx, false)
		if err != nil {
			return errors.Wrapf(err, "Failed to fetch servers")
		}
	} else {
		var server store.Server
		if errS := app.store.GetServerByName(ctx, serverName, &server); errS != nil {
			return errors.Wrapf(errS, "Failed to fetch server: %s", serverName)
		}
		servers = append(servers, server)
	}
	msg := fmt.Sprintf(`sm_csay %s`, message)
	// TODO check response
	_ = query.RCON(ctx, app.logger, servers, msg)
	app.logger.Info("Server center message sent", zap.Int64("author", author.Int64()),
		zap.String("msg", message), zap.Int("servers", len(servers)))
	return nil
}

// PSay is used to send a private message to a player
func (app *App) PSay(ctx context.Context, author steamid.SID64, target store.StringSID, message string, server *store.Server) error {
	var actualServer *store.Server
	if server != nil {
		actualServer = server
	} else {
		var playerInfo state.PlayerInfo
		// TODO check resp
		_ = app.Find(ctx, target, "", &playerInfo)
		if !playerInfo.Valid || !playerInfo.InGame {
			return consts.ErrUnknownID
		}
		actualServer = playerInfo.Server
	}
	sid, errSid := target.SID64()
	if errSid != nil {
		return errSid
	}
	msg := fmt.Sprintf(`sm_psay "#%s" "%s"`, steamid.SID64ToSID(sid), message)
	_, errExecRCON := query.ExecRCON(ctx, *actualServer, msg)
	if errExecRCON != nil {
		return errors.Errorf("Failed to exec psay command: %v", errExecRCON)
	}
	app.logger.Info("Private message sent",
		zap.Int64("author", author.Int64()),
		zap.String("message", message),
		zap.String("server", server.ServerNameShort),
		zap.Int64("target", sid.Int64()))
	return nil
}

// FilterAdd creates a new chat filter using a regex pattern
func (app *App) FilterAdd(ctx context.Context, filter *store.Filter) error {
	if errSave := app.store.SaveFilter(ctx, filter); errSave != nil {
		if errSave == store.ErrDuplicate {
			return store.ErrDuplicate
		}
		app.logger.Error("Error saving filter word", zap.Error(errSave))
		return consts.ErrInternal
	}
	filter.Init()
	wordFiltersMu.Lock()
	wordFilters = append(wordFilters, *filter)
	wordFiltersMu.Unlock()
	app.SendDiscordPayload(discordutil.Payload{
		ChannelId: config.Discord.ModLogChannelId,
		Embed: &discordgo.MessageEmbed{
			Title:       "Added new filter",
			Description: filter.Pattern,
		},
	})
	return nil
}

// FilterDel removed and existing chat filter
func (app *App) FilterDel(ctx context.Context, database store.Store, filterId int64) (bool, error) {
	var filter store.Filter
	if errGetFilter := database.GetFilterByID(ctx, filterId, &filter); errGetFilter != nil {
		return false, errGetFilter
	}
	if errDropFilter := database.DropFilter(ctx, &filter); errDropFilter != nil {
		return false, errDropFilter
	}
	wordFiltersMu.Lock()
	var valid []store.Filter
	for _, f := range wordFilters {
		if f.FilterID == filterId {
			continue
		}
		valid = append(valid, f)
	}
	wordFilters = valid
	wordFiltersMu.Unlock()
	return true, nil
}

// FilterCheck can be used to check if a phrase will match any filters
func (app *App) FilterCheck(message string) []store.Filter {
	if message == "" {
		return nil
	}
	words := strings.Split(strings.ToLower(message), " ")
	wordFiltersMu.RLock()
	defer wordFiltersMu.RUnlock()
	var found []store.Filter
	for _, filter := range wordFilters {
		for _, word := range words {
			if filter.Match(word) {
				found = append(found, filter)
			}
		}
	}
	return found
}

// findFilteredWordMatch checks to see if the body of text contains a known filtered word
// It will only return the first matched filter found.
func findFilteredWordMatch(body string) (string, *store.Filter) {
	if body == "" {
		return "", nil
	}
	words := strings.Split(strings.ToLower(body), " ")
	wordFiltersMu.RLock()
	defer wordFiltersMu.RUnlock()
	for _, filter := range wordFilters {
		for _, word := range words {
			if filter.Match(word) {
				return word, &filter
			}
		}
	}
	return "", nil
}

//// PersonBySID fetches the person from the database, updating the PlayerSummary if it out of date
//func (app *App) PersonBySID(ctx context.Context, sid steamid.SID64, person *model.Person) error {
//	if errGetPerson := app.store.GetPersonBySteamID(ctx, sid, person); errGetPerson != nil && errGetPerson != store.ErrNoResult {
//		return errGetPerson
//	}
//	if person.UpdatedOn == person.CreatedOn || time.Since(person.UpdatedOn) > 15*time.Second {
//		summary, errSummary := steamweb.PlayerSummaries(steamid.Collection{person.SteamID})
//		if errSummary != nil || len(summary) != 1 {
//			app.logger.Error("Failed to fetch updated profile", zap.Error(errSummary))
//			return nil
//		}
//		var sum = summary[0]
//		person.PlayerSummary = &sum
//		person.UpdatedOn = config.Now()
//		if errSave := app.store.SavePerson(ctx, person); errSave != nil {
//			app.logger.Error("Failed to save updated profile", zap.Error(errSummary))
//			return nil
//		}
//		if errGetPersonBySid64 := app.store.GetPersonBySteamID(ctx, sid, person); errGetPersonBySid64 != nil &&
//			errGetPersonBySid64 != store.ErrNoResult {
//			return errGetPersonBySid64
//		}
//	}
//	return nil
//}
