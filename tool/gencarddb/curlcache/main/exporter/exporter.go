package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	CacheCounter *prometheus.CounterVec
	ReqCounter *prometheus.CounterVec
)

func NewRegistry() *prometheus.Registry {
	reg := prometheus.NewRegistry()
	CacheCounter = promauto.With(reg).NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "",
			Subsystem:   "",
			Name:        "curlcache_use_cache",
			Help:        "Cache hit counter",
		},
		[]string{"isHit"},
	)

	ReqCounter = promauto.With(reg).NewCounterVec(
		prometheus.CounterOpts{
			Namespace:   "",
			Subsystem:   "",
			Name: "curlcache_request_counter",
			Help: "Http request counter",
		},
		[]string{"code"},
	)

	return reg
}
