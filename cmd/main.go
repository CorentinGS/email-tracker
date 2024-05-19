package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/bytedance/gopkg/util/gctuner"
	"github.com/corentings/email-tracker/app"
	"github.com/corentings/email-tracker/config"
)

func main() {
	gcTuning()

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

func gcTuning() {
	var limit float64 = 4 * config.GCLimit
	// Set the GC threshold to 70% of the limit
	threshold := uint64(limit * config.GCThresholdPercent)

	gctuner.Tuning(threshold)

	slog.Info(fmt.Sprintf("🔧 GC Tuning - Limit: %.2f GB, Threshold: %d bytes, GC Percent: %d, Min GC Percent: %d, Max GC Percent: %d",
		limit/(config.GCLimit),
		threshold,
		gctuner.GetGCPercent(),
		gctuner.GetMinGCPercent(),
		gctuner.GetMaxGCPercent()))

	slog.Info("✅ GC Tuning completed!")
}
