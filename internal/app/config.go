package app

import (
	"fmt"
	"github.com/leighmacdonald/bd/pkg/util"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type RunMode string

const (
	// ReleaseMode is production mode, minimal logging.
	ReleaseMode RunMode = "release"
	// DebugMode has much more logging and uses non-embedded assets.
	DebugMode RunMode = "debug"
	// TestMode is for unit tests.
	TestMode RunMode = "test"
)

// String returns the string value of the RunMode.
func (rm RunMode) String() string {
	return string(rm)
}

type Action string

const (
	Gag  Action = "gag"
	Kick Action = "kick"
	Ban  Action = "ban"
)

//// ReadConfig reads in config file and ENV variables if set.
//func ReadConfig(settings *Config, noFileOk bool) error {
//
//	if errUnmarshal := viper.Unmarshal(settings); errUnmarshal != nil {
//		return errors.Wrap(errUnmarshal, "Invalid config file format")
//	}
//
//	if strings.HasPrefix(settings.DB.DSN, "pgx://") {
//		settings.DB.DSN = strings.Replace(settings.DB.DSN, "pgx://", "postgres://", 1)
//	}
//
//	gin.SetMode(settings.General.Mode.String())
//
//	if errSteam := steamid.SetKey(settings.General.SteamKey); errSteam != nil {
//		return errors.Wrap(errSteam, "Failed to set steamid api key")
//	}
//
//	if errSteamWeb := steamweb.SetKey(settings.General.SteamKey); errSteamWeb != nil {
//		return errors.Wrap(errSteamWeb, "Failed to set steamweb api key")
//	}
//
//	_, errDuration := time.ParseDuration(settings.General.ServerStatusUpdateFreq)
//	if errDuration != nil {
//		return errors.Errorf("Failed to parse server_status_update_freq: %v", errDuration)
//	}
//
//	warnTimeoutDuration, errWarnTimeoutDuration := time.ParseDuration(settings.General.WarningTimeout)
//	if errWarnTimeoutDuration != nil {
//		return errors.Wrap(errWarnTimeoutDuration, "Failed to parse warning timeout")
//	}
//
//	settings.General.WarningTimeoutValue = warnTimeoutDuration
//
//	clientTimeoutDuration, errClientTimeoutDuration := time.ParseDuration(settings.HTTP.ClientTimeout)
//	if errClientTimeoutDuration != nil {
//		return errors.Wrap(errClientTimeoutDuration, "Failed to parse client timeout duration")
//	}
//
//	settings.HTTP.ClientTimeoutValue = clientTimeoutDuration
//
//	return nil
//}

func MustCreateLogger(runMode RunMode, logFile string) *zap.Logger {
	var loggingConfig zap.Config
	if runMode == ReleaseMode {
		loggingConfig = zap.NewProductionConfig()
		loggingConfig.DisableCaller = true
	} else {
		loggingConfig = zap.NewDevelopmentConfig()
		loggingConfig.DisableStacktrace = true
		loggingConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if logFile != "" {
		if util.Exists(logFile) {
			if err := os.Remove(logFile); err != nil {
				panic(fmt.Sprintf("Failed to remove log file: %v", err))
			}
		}

		// loggingConfig.Level.SetLevel(zap.DebugLevel)
		loggingConfig.OutputPaths = append(loggingConfig.OutputPaths, logFile)
	}

	l, errLogger := loggingConfig.Build()
	if errLogger != nil {
		panic("Failed to create log config")
	}

	return l.Named("gb")
}
