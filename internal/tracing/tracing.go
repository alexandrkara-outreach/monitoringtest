package tracing

import (
	"net/http"
	"runtime"
	"time"

	"github.com/alexandrkara-outreach/monitoringtest/internal/util"
	"github.com/honeycombio/beeline-go"
	"github.com/honeycombio/libhoney-go"
)

// AddCommonLibhoneyFields adds some dynamic generally useful fields to the current trace.
func AddCommonLibhoneyFields() {
	libhoney.AddDynamicField("meta.num_goroutines",
		func() interface{} { return runtime.NumGoroutine() })
	getAlloc := func() interface{} {
		var mem runtime.MemStats
		runtime.ReadMemStats(&mem)
		return mem.Alloc
	}
	libhoney.AddDynamicField("meta.memory_inuse", getAlloc)

	startTime := time.Now()
	libhoney.AddDynamicField("meta.process_uptime_sec", func() interface{} {
		return time.Now().Sub(startTime) / time.Second
	})
}

// ShouldSample returns whether the current trace should be included in a sample.
func ShouldSample(fields map[string]interface{}) (bool, int) {
	// Sample 10% of requests by default.
	return true, 10
}

// RecordHTTP adds a span around an incoming HTTP request.
func RecordHTTP(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, span := beeline.StartSpan(r.Context(), "http.request")
		defer span.Send()

		lw := util.NewLoggingResponseWriter(w)
		handler(lw, r.WithContext(ctx))

		span.AddField("http_uri", r.RequestURI)
		span.AddField("http_status", lw.StatusCode)
	}
}
