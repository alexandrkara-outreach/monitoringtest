package main

import (
	"github.com/alexandrkara-outreach/monitoringtest/internal/database"
	"github.com/alexandrkara-outreach/monitoringtest/internal/http"
	"github.com/alexandrkara-outreach/monitoringtest/internal/service"
)

func main() {

	db := database.NewDB()

	heavy := service.NewHeavy(db)

	c := http.NewController(heavy)

	router := http.CreateRouter(c)
	http.RunServer(router)
}
