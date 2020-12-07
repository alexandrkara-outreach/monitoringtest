package main

import (
	"log"
	"os"

	beeline "github.com/honeycombio/beeline-go"
	hnynethttp "github.com/honeycombio/beeline-go/wrappers/hnynethttp"

	"github.com/alexandrkara-outreach/monitoringtest/internal/database"
	"github.com/alexandrkara-outreach/monitoringtest/internal/http"
	"github.com/alexandrkara-outreach/monitoringtest/internal/service"
	"github.com/alexandrkara-outreach/monitoringtest/internal/stats"
	"github.com/alexandrkara-outreach/monitoringtest/internal/tracing"
)

func main() {
	// Configure Honeycomb. In reality, the key shouldn't be hardcoded.
	honeycombWriteKey := os.Getenv("HONEYCOMB_WRITEKEY")
	if honeycombWriteKey == "" {
		log.Error("You need to provide a Honeycomb key in HONEYCOMB_WRITEKEY variable")
	}
	beeline.Init(beeline.Config{
		WriteKey:    honeycombWriteKey,
		Dataset:     "MonitoringTest",
		Debug:       true,
		SamplerHook: tracing.ShouldSample,
	})
	defer beeline.Close()

	// Request that we want some common fields captured in every trace.
	tracing.AddCommonLibhoneyFields()

	// Initialize statistics collection.
	stats := stats.NewStats()

	// Create a fake connection to a DB.
	db := database.NewDB(stats)

	// Create the service we will be using to calculate the result of our "SuperEndpoint".
	heavy := service.NewHeavy(db, stats)

	// Create a HTTP controller for our service.
	c := http.NewController(heavy, stats)

	// Start serving HTTP traffic.
	router := http.CreateRouter(c, stats)
	routerWithTracing := hnynethttp.WrapHandler(router)
	http.RunServer(routerWithTracing)
}
