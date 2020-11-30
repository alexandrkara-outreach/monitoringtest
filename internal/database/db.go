package database

import (
	"context"
	"math/rand"
	"time"

	"github.com/honeycombio/beeline-go"

	"github.com/alexandrkara-outreach/monitoringtest/internal/stats"
)

// A fake database.
type DB struct {
	stats *stats.Stats
}

func NewDB(stats *stats.Stats) *DB {
	return &DB{
		stats: stats,
	}
}

// Load loads a number from a fake database.
// Normally, we would propagate the trace ID to the database.
func (db *DB) Load(ctx context.Context) int {
	ctx, span := beeline.StartSpan(ctx, "loading")
	defer span.Send()

	db.stats.Count("database.query", 1, []string{"name:load"})

	var n int

	db.stats.Measure("database.latency", []string{"name:load"}, func() {
		n = rand.Intn(100)
		time.Sleep(time.Duration(10*n) * time.Millisecond)
	})

	time.Sleep(time.Duration(10*n) * time.Millisecond)

	return n
}
