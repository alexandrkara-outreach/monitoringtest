package http

import (
	"encoding/json"
	"net/http"

	"github.com/alexandrkara-outreach/monitoringtest/internal/stats"
	"github.com/gorilla/mux"
)

func CreateRouter(c *Controler, stats *stats.Stats) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	router.HandleFunc("/super/endpoint", stats.RecordHTTP("GET-/super/endpoint", c.SuperEndpoint))

	return router
}
