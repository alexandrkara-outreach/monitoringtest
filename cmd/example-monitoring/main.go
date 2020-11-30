package main

import (
	"github.com/alexandrkara-outreach/monitoringtest/internal/database"
	"github.com/alexandrkara-outreach/monitoringtest/internal/http"
	"github.com/alexandrkara-outreach/monitoringtest/internal/monitoring"
	"github.com/alexandrkara-outreach/monitoringtest/internal/service"
)

func main() {
	stats := monitoring.NewStats()

	db := database.NewDB(stats)

	heavy := service.NewHeavy(db, stats)

	c := http.NewController(heavy)

	router := http.CreateRouter(c)
	http.RunServer(router)
}
