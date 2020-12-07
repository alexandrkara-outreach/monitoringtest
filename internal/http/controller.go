package http

import (
	"encoding/json"
	"net/http"

	"github.com/alexandrkara-outreach/monitoringtest/internal/service"
	"github.com/alexandrkara-outreach/monitoringtest/internal/stats"
)

// Controller captures the main service with additional utilities (e.g. stats).
type Controler struct {
	service *service.Heavy
	stats   *stats.Stats
}

// NewController creates a new Controller instance.
func NewController(service *service.Heavy, stats *stats.Stats) *Controler {
	return &Controler{
		service: service,
		stats:   stats,
	}
}

// SuperEndpoint does something great (calculates a Fibonacci number).
func (c *Controler) SuperEndpoint(w http.ResponseWriter, r *http.Request) {
	n, err := c.service.Compute(r.Context())

	if err != nil {
		c.stats.RecordError(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	} else {
		json.NewEncoder(w).Encode(map[string]service.Result{"computed": n})
	}
}
