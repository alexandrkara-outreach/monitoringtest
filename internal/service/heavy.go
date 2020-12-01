package service

import (
	"context"
	"errors"
	"math/big"
	"strconv"

	"github.com/honeycombio/beeline-go"

	"github.com/alexandrkara-outreach/monitoringtest/internal/database"
	"github.com/alexandrkara-outreach/monitoringtest/internal/stats"
	"github.com/alexandrkara-outreach/monitoringtest/internal/util"
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

func (h *Heavy) Compute(ctx context.Context) (Result, error) {
	var (
		r   Result
		err error
	)

	if util.Lucky(0.1) {
		return r, errors.New("service.heavy")
	}

	if util.Lucky(0.5) {
		if _, err = h.db.Query(ctx, "query_random1", 25); err != nil {
			return r, err
		}
	}

	if util.Lucky(0.3) {
		if _, err = h.db.Query(ctx, "query_random2", 50); err != nil {
			return r, err
		}
	}

	if r.Input, err = h.db.Query(ctx, "query_input", 5); err != nil {
		return r, err
	}

	r.Factorial = h.factorial(ctx, big.NewInt(int64(r.Input)))

	h.stats.Gauge("service.factorial", float64(r.Input), []string{"name:Heavy.Compute"})

	return r, nil
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
