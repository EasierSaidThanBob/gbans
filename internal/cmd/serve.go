package cmd

import (
	"context"
	"fmt"
	"github.com/leighmacdonald/gbans/internal/settings"
	"os"
	"os/signal"
	"syscall"

	"github.com/leighmacdonald/gbans/internal/app"
	"github.com/leighmacdonald/gbans/internal/store"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// serveCmd represents the serve command.
func serveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Starts the gbans service",
		Long:  `Starts the main gbans application`,
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			rootCtx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
			defer stop()

			// Start by reading static settings from config file. These are settings that cannot be changed without
			//restarting the application.
			config, errConfig := settings.SettingsReadStatic()
			if errConfig != nil {
				panic(fmt.Sprintf("Failed to read config: %v", errConfig))
			}

			rootLogger := app.MustCreateLogger(config.GeneralRunMode, config.LogFile)
			defer func() {
				if config.LogFile != "" {
					_ = rootLogger.Sync()
				}
			}()

			database := store.New(rootLogger, config.DatabaseDSN, config.DatabaseAutoMigrate, config.DatabaseLogQueries)
			if errConnect := database.Connect(rootCtx); errConnect != nil {
				rootLogger.Fatal("Cannot initialize database", zap.Error(errConnect))
			}

			defer func() {
				if errClose := database.Close(); errClose != nil {
					rootLogger.Error("Failed to close database cleanly")
				}
			}()

			dbSettings, errSettings := database.Settings(ctx)
			if errSettings != nil {
				panic(errSettings)
			}

			if err := settings.ReadDB(database, dbSettings); err != nil {
				panic(err)
			}

			application := app.New(&config, database, rootLogger)

			if errInit := application.Init(rootCtx); errInit != nil {
				rootLogger.Fatal("Failed to init app", zap.Error(errInit))
			}

			if errWebStart := application.Listen(rootCtx); errWebStart != nil {
				rootLogger.Error("Web returned error", zap.Error(errWebStart))
			}

			<-rootCtx.Done()
		},
	}
}
