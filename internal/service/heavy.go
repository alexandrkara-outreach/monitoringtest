package service

import (
	"context"
	"math/big"
	"strconv"

	"github.com/honeycombio/beeline-go"

	"github.com/alexandrkara-outreach/monitoringtest/internal/database"
	"github.com/alexandrkara-outreach/monitoringtest/internal/monitoring"
)

type Heavy struct {
	db    *database.DB
	stats *monitoring.Stats
}

type Result struct {
	Input     int      `json:"input"`
	Factorial *big.Int `json:"factorial"`
}

func NewHeavy(db *database.DB, stats *monitoring.Stats) *Heavy {
	return &Heavy{
		db:    db,
		stats: stats,
	}
}

func (h *Heavy) Compute(ctx context.Context) Result {
	var r Result
	r.Input = h.db.Load(ctx)
	r.Factorial = h.factorial(ctx, big.NewInt(int64(r.Input)))
	return r
}

func (h *Heavy) factorial(ctx context.Context, x *big.Int) *big.Int {
	ctx, span := beeline.StartSpan(ctx, "factorial")
	defer span.Send()

	beeline.AddField(ctx, "value", strconv.FormatInt(x.Int64(), 10))

	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}

	return n.Mul(x, h.factorial(ctx, n.Sub(x, n)))
}
