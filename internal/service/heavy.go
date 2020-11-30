package service

import (
	"math/big"

	"github.com/alexandrkara-outreach/monitoringtest/internal/database"
	"github.com/alexandrkara-outreach/monitoringtest/internal/stats"
)

type Heavy struct {
	db    *database.DB
	stats *stats.Stats
}

type Result struct {
	Input     int      `json:"input"`
	Factorial *big.Int `json:"factorial"`
}

func NewHeavy(db *database.DB, stats *stats.Stats) *Heavy {
	return &Heavy{
		db:    db,
		stats: stats,
	}
}

func (h *Heavy) Compute() Result {
	var r Result
	r.Input = h.db.Load()
	r.Factorial = h.factorial(big.NewInt(int64(r.Input)))

	h.stats.Gauge("service.factorial", float64(r.Input), []string{"name:Heavy.Compute"})

	return r
}

func (h *Heavy) factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, h.factorial(n.Sub(x, n)))
}
