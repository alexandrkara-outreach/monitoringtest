package database

import (
	"context"
	"math/rand"
	"time"

	"github.com/honeycombio/beeline-go"

	"github.com/alexandrkara-outreach/monitoringtest/internal/monitoring"
)

// A fake database.
type DB struct {
	stats *monitoring.Stats
}

func NewDB(stats *monitoring.Stats) *DB {
	return &DB{
		stats: stats,
	}
}

// Load loads a number from a fake database.
// Normally, we would propagate the trace ID to the database.
func (db *DB) Load(ctx context.Context) int {
	ctx, span := beeline.StartSpan(ctx, "loading")
	defer span.Send()

	n := rand.Intn(100)
	time.Sleep(time.Duration(10*n) * time.Millisecond)
	return n
}
