package database

import (
	"math/rand"
	"time"

	"github.com/alexandrkara-outreach/monitoringtest/internal/monitoring"
)

type DB struct {
	stats *monitoring.Stats
}

func NewDB(stats *monitoring.Stats) *DB {
	return &DB{
		stats: stats,
	}
}

func (db *DB) Load() int {
	n := rand.Intn(100)
	time.Sleep(time.Duration(10*n) * time.Millisecond)
	return n
}
