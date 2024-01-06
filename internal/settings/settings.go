package settings

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/leighmacdonald/gbans/internal/app"
	"github.com/leighmacdonald/gbans/internal/store"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
	"time"

	"github.com/leighmacdonald/steamid/v3/steamid"
)

type Static struct {
	GeneralOwner    steamid.SID64 `mapstructure:"general_owner"`
	GeneralSteamKey string        `config_key:"general_steam_key"`
	GeneralRunMode  app.RunMode   `config_key:"general_run_mode"`

	DatabaseDSN         string `config_key:"database_dsn"`
	DatabaseLogQueries  bool   `config_key:"database_log_queries"`
	DatabaseAutoMigrate bool   `mapstructure:"database_auto_migrate"`

	S3AccessKey   string `config_key:"s3_access_key"`
	S3SecretKey   string `config_key:"s3_secret_key"`
	S3Endpoint    string `config_key:"s3_endpoint"`
	S3ExternalURL string `config_key:"s3_external_url"`
	S3SSL         bool   `config_key:"s3_ssl"`
	S3Region      string `config_key:"s3_region"`
	S3BucketMedia string `config_key:"s3_bucket_media"`
	S3BucketDemo  string `config_key:"s3_bucket_demo"`

	HTTPListenHost    string        `config_key:"http_host"`
	HTTPListenPort    uint16        `config_key:"http_port"`
	HTTPStaticPath    string        `config_key:"http_static_path"`
	HTTPClientTimeout time.Duration `config_key:"http_client_timeout"`
	HTTPCookieKey     string        `config_key:"http_cookie_key"`
	HTTPCorsOrigins   []string      `config_key:"http_cors_origins"`
	HTTPExternalURL   string        `config_key:"http_external_url"`
}

type Dynamic struct {
	GeneralSiteName         string        `config_key:"general_site_name"`
	GeneralStatusUpdateFreq time.Duration `config_key:"general_server_status_update_freq"`
	GeneralDemoCountLimit   int           `config_key:"general_demo_count_limit"`

	GeneralPatreonEnabled bool `config_key:"patreon_patreon_enabled"`

	DebugUpdateSRCDSLogSecrets   bool   `config_key:"debug_update_srcds_log_secrets"`
	DebugRCONAddress             string `config_key:"debug_add_rcon_log_address"`
	DebugSkipOpenIDValidation    bool   `config_key:"debug_skip_open_id_validation"`
	DebugWriteUnhandledLogEvents bool   `config_key:"debug_write_unhandled_log_events"`

	WordFilterEnabled      bool `config_key:"word_filter_enabled"`
	WordFilterPingDiscord  bool `config_key:"word_filter_ping_discord"`
	WordFilterDry          bool `config_key:"word_filter_dry_run"`
	WordFilterWarningLimit int  `config_key:"word_filter_warning_limit"`

	DiscordEnabled                 bool   `config_key:"discord_enabled"`
	DiscordLinkID                  string `config_key:"discord_link_id"`
	DiscordAppID                   string `config_key:"discord_app_id"`
	DiscordAppSecret               string `config_key:"discord_app_secret"`
	DiscordGuildID                 string `config_key:"discord_guild_id"`
	DiscordToken                   string `config_key:"discord_token"`
	DiscordUnregisterOnStart       bool   `config_key:"discord_unregister_on_start"`
	DiscordModLogChannelID         string `config_key:"discord_mod_log_channel_id"`
	DiscordPublicLogChannelEnable  bool   `config_key:"discord_public_log_channel_enable"`
	DiscordPublicLogChannelID      string `config_key:"discord_public_log_channel_id"`
	DiscordPublicMatchLogChannelID string `config_key:"discord_public_match_log_channel_id"`

	SRCDSLogListenAddr   string `config_key:"srcds_list_listen_addr"`
	SRCDSLogExternalHost string `config_key:"srcds_external_host"`

	LogLevel         string `mapstructure:"log_level"`
	LogFile          string `mapstructure:"log_file"`
	LogReportCaller  bool   `mapstructure:"log_report_caller"`
	LogFullTimestamp bool   `mapstructure:"log_full_timestamp"`

	IP2LocationEnabled      string `config_key:"ip2location_enabled"`
	IP2LocationToken        string `config_key:"ip2location_token"`
	IP2LocationCachePath    string `config_key:"ip2location_cache_path"`
	IP2LocationASNEnabled   string `config_key:"ip2location_asn_enabled"`
	IP2LocationIPEnabled    string `config_key:"ip2location_ip_enabled"`
	IP2LocationProxyEnabled string `config_key:"ip2location_proxy_enabled"`
}

type Settings struct {
	Static
	Dynamic
}

func (c Settings) Addr() string {
	return fmt.Sprintf("%s:%d", c.HTTPListenHost, c.HTTPListenPort)
}

func settingsInit(noFileOk bool) error {
	if errReadConfig := viper.ReadInConfig(); errReadConfig != nil && !noFileOk {
		return errors.Wrapf(errReadConfig, "Failed to read config file")
	}

	return nil
}

// SettingsReadStatic will load the static settings from the config file. You must call SettingsReadDB to load
// the remaining dynamic settings.
func SettingsReadStatic() (Settings, error) {
	var s Settings
	if errUnmarshal := viper.Unmarshal(&s); errUnmarshal != nil {
		return s, errors.Wrap(errUnmarshal, "Invalid config file format")
	}

	if strings.HasPrefix(s.DatabaseDSN, "pgx://") {
		s.DatabaseDSN = strings.Replace(s.DatabaseDSN, "pgx://", "postgres://", 1)
	}

	return s, nil
}

func ReadDB(db *store.Store, s map[string]string) error {
	return nil
}

func Write(ctx context.Context, conn pgx.Conn, settings Settings) error {
	return nil
}

func init() {
	if home, errHomeDir := homedir.Dir(); errHomeDir != nil {
		viper.AddConfigPath(home)
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("gbans")
	viper.SetConfigType("yml")
	viper.SetEnvPrefix("gbans")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	defaultConfig := map[string]any{
		"general.site_name":                        "gbans",
		"general.steam_key":                        "",
		"general.mode":                             "release",
		"general.owner":                            76561198044052046,
		"general.warning_timeout":                  "72h",
		"general.warning_limit":                    2,
		"general.warning_exceeded_action":          "gag",
		"general.warning_exceeded_duration":        "168h",
		"general.use_utc":                          true,
		"general.server_status_update_freq":        "60s",
		"general.master_server_status_update_freq": "1m",
		"general.external_url":                     "http://gbans.localhost:6006",
		"general.banned_steam_group_ids":           []steamid.GID{},
		"general.banned_server_addresses":          []string{},
		"general.demo_cleanup_enabled":             true,
		"general.demo_count_limit":                 10000,
		"patreon.enabled":                          false,
		"patreon.client_id":                        "",
		"patreon.client_secret":                    "",
		"patreon.creator_access_token":             "",
		"patreon.creator_refresh_token":            "",
		"http.host":                                "127.0.0.1",
		"http.port":                                6006,
		"http.tls":                                 false,
		"http.tls_auto":                            false,
		"http.static_path":                         "frontend/dist",
		"http.cookie_key":                          store.SecureRandomString(32),
		"http.client_timeout":                      "10s",
		"debug.update_srcds_log_secrets":           true,
		"debug.skip_open_id_validation":            false,
		"debug.write_unhandled_log_events":         false,
		"filter.enabled":                           false,
		"filter.dry":                               true,
		"filter.ping_discord":                      false,
		"discord.enabled":                          false,
		"discord.app_id":                           0,
		"discord.app_secret":                       "",
		"discord.token":                            "",
		"discord.link_id":                          "",
		"discord.perms":                            125958,
		"discord.guild_id":                         "",
		"discord.public_log_channel_enable":        false,
		"discord.public_log_channel_id":            "",
		"discord.public_match_log_channel_id":      "",
		"discord.log_channel_id":                   "",
		"discord.mod_ping_role_id":                 "",
		"discord.unregister_on_start":              false,
		"ip2location.enabled":                      false,
		"ip2location.token":                        "",
		"ip2location.asn_enabled":                  false,
		"ip2location.ip_enabled":                   false,
		"ip2location.proxy_enabled":                false,
		"log.level":                                "info",
		"log.report_caller":                        false,
		"log.full_timestamp":                       false,
		"log.srcds_log_addr":                       ":27115",
		"database.dsn":                             "postgresql://gbans:gbans@localhost/gbans",
		"database.auto_migrate":                    true,
		"database.log_queries":                     false,
		"s3.enabled":                               false,
		"s3.access_key":                            "",
		"s3.secret_key":                            "",
		"s3.endpoint":                              "localhost:9001",
		"s3.ssl":                                   false,
		"s3.region":                                "",
		"s3.bucket_media":                          "media",
		"s3.bucket_demo":                           "demos",
	}

	for configKey, value := range defaultConfig {
		viper.SetDefault(configKey, value)
	}
}
