package http

import (
	"encoding/json"
	"net/http"

	"github.com/alexandrkara-outreach/monitoringtest/internal/stats"
	"github.com/alexandrkara-outreach/monitoringtest/internal/tracing"
	"github.com/gorilla/mux"
)

// CreateRouter creates HTTP endpoints for health checking and for "SuperEndpoint".
func CreateRouter(c *Controler, stats *stats.Stats) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	// Apply metric and tracing "middlewares"
	superEndpoint := c.SuperEndpoint
	superEndpoint = stats.RecordHTTP("GET-/super/endpoint", superEndpoint)
	superEndpoint = tracing.RecordHTTP(superEndpoint)
	router.HandleFunc("/super/endpoint", superEndpoint)

	return router
}
