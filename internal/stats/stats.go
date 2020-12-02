package stats

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/alexandrkara-outreach/monitoringtest/internal/util"
)

type Stats struct {
	client *statsd.Client
}

func NewStats() *Stats {
	statsd, err := statsd.New("127.0.0.1:8125")

	statsd.SkipErrors = false
	statsd.Namespace = "example-monitoring."

	if err != nil {
		panic(err)
	}
	return &Stats{
		client: statsd,
	}
}

func (s *Stats) RecordHTTP(name string, handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		lw := util.NewLoggingResponseWriter(w)
		handler(lw, r)
		status := strconv.Itoa(lw.StatusCode/100) + "XX"
		s.Timing("http.request", time.Since(t), []string{"endpoint:" + name, "status:" + status})
		log.Printf("msg:request served status:%s\n", status)
	}
}

func (s *Stats) Measure(name string, tags []string, callback func()) {
	t := time.Now()
	callback()
	s.Timing(name, time.Since(t), tags)
}

func (s *Stats) Gauge(name string, value float64, tags []string) {
	err := s.client.Gauge(name, value, tags, 1.0)
	if err != nil {
		panic(err)
	}
}

func (s *Stats) Count(name string, value int64, tags []string) {
	err := s.client.Count(name, value, tags, 1.0)
	if err != nil {
		panic(err)
	}
}

func (s *Stats) Histogram(name string, value float64, tags []string) {
	err := s.client.Histogram(name, value, tags, 1.0)
	if err != nil {
		panic(err)
	}
}

func (s *Stats) Timing(name string, value time.Duration, tags []string) {
	err := s.client.Timing(name, value, tags, 1.0)
	if err != nil {
		panic(err)
	}
}

func (s *Stats) RecordError(err error) {
	s.Count("errors", 1, []string{"error:" + err.Error()})
}
