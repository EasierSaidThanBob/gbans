package domain

import (
	"errors"
)

var (
	ErrMessageContext           = errors.New("could not fetch message context")
	ErrFailedWeapon             = errors.New("failed to save weapon")
	ErrSaveChanges              = errors.New("cannot save changes")
	ErrDiscordAlreadyLinked     = errors.New("discord account is already linked")
	ErrFailedFetchBan           = errors.New("failed to get existing ban")
	ErrSaveBan                  = errors.New("failed to save ban")
	ErrGetBanReport             = errors.New("failed to get ban report")
	ErrReportStateUpdate        = errors.New("failed to update report state")
	ErrFetchPerson              = errors.New("failed to fetch/create person")
	ErrFetchSource              = errors.New("failed to fetch source player")
	ErrFetchTarget              = errors.New("failed to fetch target player")
	ErrSaveBanGroup             = errors.New("failed to save ban group")
	ErrGroupValidate            = errors.New("failed to validate group")
	ErrParseASN                 = errors.New("failed to parse asn")
	ErrFetchASNBan              = errors.New("failed to fetch asn ban")
	ErrDropASNBan               = errors.New("failed to drop existing asn ban")
	ErrUnknownASN               = errors.New("no networks found for asn")
	ErrFetchASN                 = errors.New("failed to fetch asn networks")
	ErrASNNoRecords             = errors.New("no records found for asn")
	ErrOwnerInvalid             = errors.New("configured owner steamid is invalid")
	ErrCreateAdmin              = errors.New("failed to create admin user")
	ErrSetupAdmin               = errors.New("failed to setup admin user")
	ErrSetupNews                = errors.New("failed to create sample news entry")
	ErrSetupServer              = errors.New("failed to create sample server entry")
	ErrSetupWiki                = errors.New("failed to create sample wiki entry")
	ErrSetupWeapons             = errors.New("failed to setup weapon maps")
	ErrLoadFilters              = errors.New("failed to load chat filters")
	ErrInvalidFilterRegex       = errors.New("could not parse filter regex")
	ErrInvalidPattern           = errors.New("invalid pattern")
	ErrInvalidFilterID          = errors.New("invalid fiter ID")
	ErrInitNetBlocks            = errors.New("failed to load net blocks")
	ErrInitNetWhitelist         = errors.New("failed to load net block whitelists")
	ErrHTTPServer               = errors.New("http listener returned unexpected error")
	ErrNotificationSteamIDs     = errors.New("failed to load notification steam ids")
	ErrNotificationPeople       = errors.New("failed to load notified people")
	ErrSteamAPIArgLimit         = errors.New("steam api support a max of 100 steam ids")
	ErrFetchSteamBans           = errors.New("failed to fetch ban status from steam api")
	ErrSteamAPISummaries        = errors.New("failed to fetch player summaries")
	ErrSteamAPI                 = errors.New("steam api requests have errors")
	ErrUpdatePerson             = errors.New("failed to save updated person profile")
	ErrCommandFailed            = errors.New("command failed")
	ErrDiscordCreate            = errors.New("failed to connect to discord")
	ErrDiscordOpen              = errors.New("failed to open discord connection")
	ErrDuplicateCommand         = errors.New("duplicate command registration")
	ErrDiscordMessageSen        = errors.New("failed to send discord message")
	ErrDiscordOverwriteCommands = errors.New("failed to bulk overwrite discord commands")
	ErrInsufficientPlayers      = errors.New("insufficient Match players")
	ErrIncompleteMatch          = errors.New("insufficient match data")
	ErrSaveMatch                = errors.New("could not save match results")
	ErrLoadMatch                = errors.New("could not load match results")
	ErrLoadServer               = errors.New("failed to load match server")
	ErrPatreonFetchCampaign     = errors.New("failed to fetch patreon campaign")
	ErrPatreonInvalidCampaign   = errors.New("patreon campaign does not exist")
	ErrPatreonFetchPledges      = errors.New("failed to fetch patreon pledges")
	ErrRegisterCommand          = errors.New("failed to register discord command")
	ErrBanDoesNotExist          = errors.New("ban does not exist")
	ErrSteamUnset               = errors.New("must set steam id. see /set_steam command")
	ErrFetchClassStats          = errors.New("failed to fetch class stats")
	ErrFetchWeaponStats         = errors.New("failed to fetch weapon stats")
	ErrFetchKillstreakStats     = errors.New("failed to fetch killstreak stats")
	ErrFetchMedicStats          = errors.New("failed to fetch medic stats")
	ErrGetServer                = errors.New("failed to get server")
	ErrReasonInvalid            = errors.New("invalid reason")
	ErrDuplicateBan             = errors.New("duplicate ban")
	ErrCIDRMissing              = errors.New("cidr invalid or missing")
	ErrRowResults               = errors.New("resulting rows contain error")
	ErrTxStart                  = errors.New("could not start transaction")
	ErrTxCommit                 = errors.New("failed to commit tx changes")
	ErrTxRollback               = errors.New("could not rollback transaction")
	ErrPoolFailed               = errors.New("could not create store pool")
	ErrUUIDGen                  = errors.New("could not generate uuid")
	ErrCreateQuery              = errors.New("failed to generate query")
	ErrCountQuery               = errors.New("failed to get count result")
	ErrTooShort                 = errors.New("value too short")
	ErrInvalidParameter         = errors.New("invalid parameter format")
	ErrPermissionDenied         = errors.New("permission denied")
	ErrBadRequest               = errors.New("invalid request")
	ErrInternal                 = errors.New("internal server error")
	ErrParamKeyMissing          = errors.New("param key not found")
	ErrParamParse               = errors.New("failed to parse param value")
	ErrParamInvalid             = errors.New("param value invalid")
	ErrScanResult               = errors.New("failed to scan result")
	ErrUnknownServerID          = errors.New("unknown server id")
	ErrSelfReport               = errors.New("cannot self report")
	ErrUUIDCreate               = errors.New("failed to generate new uuid")
	ErrReportExists             = errors.New("cannot create report while existing report open")
	ErrAssetCreateFailed        = errors.New("failed to create asset")
	ErrAssetPut                 = errors.New("unable to create asset on remote store")
	ErrAssetGet                 = errors.New("failed to get asset from client")
	ErrAssetSave                = errors.New("failed to save asset metadata")
	ErrInvalidFormat            = errors.New("invalid format")
	ErrDuplicateMediaName       = errors.New("duplicate media name")
	ErrSaveMedia                = errors.New("could not save media")
	ErrFetchMedia               = errors.New("failed to fetch media asset")
	ErrEmptyToken               = errors.New("invalid Access token decoded")
	ErrContestLoadEntries       = errors.New("failed to load existing contest entries")
	ErrContestMaxEntries        = errors.New("entries count exceed max_submission limits")
	ErrEntryCreate              = errors.New("failed to create new contest entry")
	ErrEntrySave                = errors.New("failed to save contest entry")
	ErrThreadLocked             = errors.New("thread is locked")
	ErrCreateToken              = errors.New("failed to create new Access token")
	ErrRefreshToken             = errors.New("failed to create new Refresh token")
	ErrClientIP                 = errors.New("failed to parse IP")
	ErrSaveToken                = errors.New("failed to save new createRefresh token")
	ErrSignToken                = errors.New("failed create signed string")
	ErrSignJWT                  = errors.New("failed create signed JWT string")
	ErrAuthHeader               = errors.New("failed to bind auth header")
	ErrMalformedAuthHeader      = errors.New("invalid auth header format")
	ErrCookieKeyMissing         = errors.New("cookie key missing, cannot generate token")
	ErrInvalidContestID         = errors.New("invalid contest id")
	ErrInvalidDescription       = errors.New("invalid description, cannot be empty")
	ErrTitleEmpty               = errors.New("title cannot be empty")
	ErrDescriptionEmpty         = errors.New("description cannot be empty")
	ErrEndDateBefore            = errors.New("end date comes before start date")
	ErrInvalidThread            = errors.New("invalid thread id")
	ErrPersonSource             = errors.New("failed to load source person")
	ErrPersonTarget             = errors.New("failed to load target person")
	ErrGetBan                   = errors.New("failed to load existing ban")
	ErrScanASN                  = errors.New("failed to scan asn result")
	ErrCloseBatch               = errors.New("failed to close batch request")
	ErrDecodeDuration           = errors.New("failed to decode duration")
	ErrReadConfig               = errors.New("failed to read config file")
	ErrFormatConfig             = errors.New("config file format invalid")
	ErrSteamAPIKey              = errors.New("failed to set steam api key")
	ErrUnbanFailed              = errors.New("failed to perform unban")
	ErrStateUnchanged           = errors.New("state must be different than previous")
	ErrInvalidRegex             = errors.New("invalid regex format")
	ErrInvalidWeight            = errors.New("invalid weight value")
	ErrReportCountQuery         = errors.New("failed to get reports count for demo")
	ErrMatchQuery               = errors.New("failed to load match")
	ErrQueryPlayers             = errors.New("failed to query match players")
	ErrQueryMatch               = errors.New("failed to query match")
	ErrChatQuery                = errors.New("failed to query chat history")
	ErrGetPlayerClasses         = errors.New("failed to fetch player class stats")
	ErrGetMedicStats            = errors.New("failed to fetch medic class stats")
	ErrSaveMedicStats           = errors.New("failed to save medic stats")
	ErrSavePlayerStats          = errors.New("failed to save player stats")
	ErrSaveWeaponStats          = errors.New("failed to save weapon stats")
	ErrSaveClassStats           = errors.New("failed to save class stats")
	ErrSaveKillstreakStats      = errors.New("failed to save killstreak stats")
	ErrGetWeaponStats           = errors.New("failed to fetch match weapon stats")
	ErrGetPlayerKillstreaks     = errors.New("failed to fetch player killstreak stats")
	ErrGetPerson                = errors.New("failed to fetch person result")
	ErrInvalidIP                = errors.New("invalid ip, could not parse")
	ErrAuthentication           = errors.New("auth invalid")
	ErrExpired                  = errors.New("expired")
	ErrInvalidSID               = errors.New("invalid steamid")
	ErrSourceID                 = errors.New("invalid source steam id")
	ErrTargetID                 = errors.New("invalid target steam id")
	ErrPlayerNotFound           = errors.New("could not find player")
	ErrInvalidTeam              = errors.New("invalid team")
	ErrUnknownID                = errors.New("could not find matching server/player/steamid")
	ErrInvalidAuthorSID         = errors.New("invalid author steam id")
	ErrInvalidTargetSID         = errors.New("invalid target steam id")
	ErrNotFound                 = errors.New("entity not found")
	ErrNoResult                 = errors.New("no results found")
	ErrDuplicate                = errors.New("entity already exists")
	ErrUnknownServer            = errors.New("unknown server")
	ErrVoteDeleted              = errors.New("vote deleted")
	ErrCreateRequest            = errors.New("failed to create new request")
	ErrRequestPerform           = errors.New("could not perform http request")
	ErrRequestInvalidCode       = errors.New("invalid response code returned from request")
	ErrRequestDecode            = errors.New("failed to decode http response")
	ErrResponseBody             = errors.New("failed to read response body")
	ErrQueryPatreon             = errors.New("failed to query patreon auth token")
	ErrMimeTypeNotAllowed       = errors.New("mimetype is not allowed")
	ErrInvalidMediaMimeType     = errors.New("detected mimetype different than type provided")
	ErrInitClient               = errors.New("failed to initialize client")
	ErrBucketCheck              = errors.New("could not determine if bucket exists")
	ErrBucketCreate             = errors.New("could not create new bucket")
	ErrPolicy                   = errors.New("failed to set bucket policy")
	ErrWriteObject              = errors.New("failed to write object")
	ErrDeleteObject             = errors.New("failed to delete object")
	ErrReadContent              = errors.New("failed to read content")
	ErrFailedToBan              = errors.New("failed to create warning ban")
	ErrWarnActionApply          = errors.New("failed to apply warning action")
	ErrStaticPathError          = errors.New("could not load static path")
	ErrDataUpdate               = errors.New("data update failed")
	ErrValidateURL              = errors.New("could not validate url")
	ErrTempDir                  = errors.New("failed to create temp dir")
	ErrOpenFile                 = errors.New("could not open output file")
	ErrWriteDemo                = errors.New("could not write demo to disk for reading")
	ErrFrontendRoutes           = errors.New("failed to initialize frontend asset routes")
)
