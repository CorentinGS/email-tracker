package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/corentings/email-tracker/app"
	"github.com/corentings/email-tracker/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("❌ Error loading config: %s", err.Error())
	}

	if err = app.Run(cfg); err != nil {
		log.Fatalf("❌ Error running app: %s", err.Error())
	}

	slog.Info("✅ Server started successfully!")

	os.Exit(0)
}
