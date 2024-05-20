package app

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/corentings/email-tracker/config"
	"github.com/corentings/email-tracker/pkg/crypto"
	"github.com/corentings/email-tracker/pkg/jwt"
	"github.com/corentings/email-tracker/pkg/logger"
	"github.com/corentings/email-tracker/pkg/postgres"
)

func Run(cfg *config.Config) error {
	logger.GetLogger().SetLogLevel(cfg.Log.Level).CreateGlobalHandler()

	if logger.GetLogger().GetLogLevel() == slog.LevelDebug {
		slog.Debug("🔧 Debug mode enabled")
	}

	if err := postgres.New(cfg.PG.URL, postgres.WithMaxPoolSize(cfg.PG.PoolMax)); err != nil {
		log.Fatalf("❌ Error connecting to database: %s", err.Error())
	}

	// Parse the keys
	km, err := crypto.NewKeyManager()
	if err != nil {
		slog.Error("error creating key manager", slog.Any("error", err))
		return err
	}

	// Create the JWT instance
	jwtInstance := jwt.NewJWTInstance(cfg.JWT.HeaderLen, cfg.JWT.Expiration,
		km.GetPublicKey(), km.GetPrivateKey())

	jwt.GetJwtInstance().SetJwt(jwtInstance)

	slog.Info("starting server 🚀")

	e := NewEcho(cfg)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err = e.Start(":" + cfg.Port); err != nil {
			slog.Error("error starting server", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	const shutdownTimeout = 10 * time.Second

	<-ctx.Done()
	_, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err = e.Shutdown(ctx); err != nil {
		slog.Error("error shutting down server", slog.Any("error", err))
	}

	slog.Info("server stopped ✋")

	return nil
}
