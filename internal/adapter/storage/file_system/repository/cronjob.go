package repository

import (
	"events_app/internal/core/util"
	"log/slog"
	"time"

	"github.com/robfig/cron/v3"
)

// Remove old events from CurrentEventsState
func RemoveOldData() {
	fiveHoursAgo := time.Now().Add(-5 * time.Hour)

	for i, e := range CurrentEventsState {
		if e.StartsAt.Before(fiveHoursAgo) {
			CurrentEventsState = util.RemoveEventFromSlice(CurrentEventsState, i)
		}
	}

}

// Start the cron scheduler
func StartCronjobs() {
	cronjob := cron.New()

	// Add a cron job
	_, err := cronjob.AddFunc("*/10 * * * *", RemoveOldData)
	if err != nil {
		slog.Error("Error adding cron job: %v", "error", err)
	}

	cronjob.Start()
}
