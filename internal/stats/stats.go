package stats

import (
	"net/http"
	"time"

	"github.com/DataDog/datadog-go/statsd"
)

type Stats struct {
	client *statsd.Client
}

func NewStats() *Stats {
	statsd, err := statsd.New("127.0.0.1:8125")
	if err != nil {
		panic(err)
	}
	return &Stats{
		client: statsd,
	}
}

func (s *Stats) RecordHTTP(name string, handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Measure("http.request", []string{"endpoint:" + name}, func() {
			handler(w, r)
		})
	}
}

func (s *Stats) Measure(name string, tags []string, callback func()) {
	t := time.Now()
	callback()
	s.Histogram(name, float64(time.Since(t).Milliseconds()), tags)
}

func (s *Stats) Gauge(name string, value float64, tags []string) {
	s.client.Gauge(name, value, tags, 1.0)
}

func (s *Stats) Count(name string, value int64, tags []string) {
	s.client.Count(name, value, tags, 1.0)
}

func (s *Stats) Histogram(name string, value float64, tags []string) {
	s.client.Histogram(name, value, tags, 1.0)
}
