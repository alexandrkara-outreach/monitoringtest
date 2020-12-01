package main

import (
	"os"

	beeline "github.com/honeycombio/beeline-go"
	hnynethttp "github.com/honeycombio/beeline-go/wrappers/hnynethttp"

	"github.com/alexandrkara-outreach/monitoringtest/internal/database"
	"github.com/alexandrkara-outreach/monitoringtest/internal/http"
	"github.com/alexandrkara-outreach/monitoringtest/internal/service"
	"github.com/alexandrkara-outreach/monitoringtest/internal/tracing"
	"github.com/alexandrkara-outreach/monitoringtest/internal/stats"
)

func main() {
	// Configure Honeycomb. In reality, the key shouldn't be hardcoded.
	honeycombWriteKey := os.Getenv("HONEYCOMB_WRITEKEY")
	if honeycombWriteKey == "" {
		honeycombWriteKey = "44f82e3dcff4d8bd02ce2271a6be7c03"
	}
	beeline.Init(beeline.Config{
		WriteKey:    honeycombWriteKey,
		Dataset:     "MonitoringTest",
		Debug:       true,
		SamplerHook: tracing.ShouldSample,
	})
	defer beeline.Close()

	tracing.AddCommonLibhoneyFields()

	stats := stats.NewStats()

	db := database.NewDB(stats)

	heavy := service.NewHeavy(db, stats)

	c := http.NewController(heavy, stats)

	router := http.CreateRouter(c, stats)
	routerWithTracing := hnynethttp.WrapHandler(router)
	http.RunServer(routerWithTracing)
}
