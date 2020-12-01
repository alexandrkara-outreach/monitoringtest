package main

import (
	"github.com/alexandrkara-outreach/monitoringtest/internal/database"
	"github.com/alexandrkara-outreach/monitoringtest/internal/http"
	"github.com/alexandrkara-outreach/monitoringtest/internal/service"
	"github.com/alexandrkara-outreach/monitoringtest/internal/stats"
)

func main() {
	stats := stats.NewStats()

	db := database.NewDB(stats)

	heavy := service.NewHeavy(db, stats)

	c := http.NewController(heavy, stats)

	router := http.CreateRouter(c, stats)
	http.RunServer(router)
}
