package database

import (
	"math/rand"
	"time"

	"github.com/alexandrkara-outreach/monitoringtest/internal/stats"
)

type DB struct {
	stats *stats.Stats
}

func NewDB(stats *stats.Stats) *DB {
	return &DB{
		stats: stats,
	}
}

func (db *DB) Load() int {
	db.stats.Count("database.query", 1, []string{"name:load"})

	var n int

	db.stats.Measure("database.latency", []string{"name:load"}, func() {
		n = rand.Intn(100)
		time.Sleep(time.Duration(10*n) * time.Millisecond)
	})

	return n
}
