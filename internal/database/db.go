package database

import (
	"errors"
	"math/rand"
	"time"

	"github.com/alexandrkara-outreach/monitoringtest/internal/stats"
	"github.com/alexandrkara-outreach/monitoringtest/internal/util"
)

var ErrChaosError = errors.New("Chaos error")

type DB struct {
	stats *stats.Stats
}

func NewDB(stats *stats.Stats) *DB {
	return &DB{
		stats: stats,
	}
}

func (db *DB) Query(name string, n int) (int, error) {
	if util.Lucky(0.1) {
		return 0, errors.New("query." + name)
	}

	db.stats.Count("database.query", 1, []string{"name:" + name})

	db.stats.Measure("database.latency", []string{"name:" + name}, func() {
		n = rand.Intn(n)
		time.Sleep(time.Duration(10*n) * time.Millisecond)
	})

	return n, nil
}
