package database

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/honeycombio/beeline-go"

	"github.com/alexandrkara-outreach/monitoringtest/internal/stats"
	"github.com/alexandrkara-outreach/monitoringtest/internal/util"
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

// Query loads a number from a fake database.
// Normally, we would propagate the trace ID to the database.
func (db *DB) Query(ctx context.Context, name string, n int) (r int, err error) {
	ctx, span := beeline.StartSpan(ctx, "query")
	defer func() {
		if err != nil {
			span.AddField("error", err.Error())
		}
		span.Send()
	}()

	if util.Lucky(0.1) {
		return 0, errors.New("query." + name)
	}

	span.AddField("query_name", name)

	db.stats.Count("database.query", 1, []string{"name:load"})

	db.stats.Count("database.query", 1, []string{"name:" + name})

	db.stats.Measure("database.latency", []string{"name:" + name}, func() {
		n = rand.Intn(n)
		time.Sleep(time.Duration(10*n) * time.Millisecond)
	})

	return n, nil
}
