package tcl_client

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go/log"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/uber-go/tally/v4"
	"github.com/uber-go/tally/v4/prometheus"
	"go.temporal.io/sdk/client"
	tallyhandler "go.temporal.io/sdk/contrib/tally"
)

const TimerType = prometheus.HistogramTimerType

func NewPrometheusScope(ctx context.Context, p prom.Registerer, prefix string) tally.Scope {
	scopeOpts := tally.ScopeOptions{
		CachedReporter: prometheus.NewReporter(prometheus.Options{
			Registerer:       p,
			DefaultTimerType: TimerType,
			OnRegisterError: func(err error) {
				log.Error(err)
			},
		}),
		Separator: prometheus.DefaultSeparator,
		Prefix:    prefix,
	}

	scope, _ := tally.NewRootScope(scopeOpts, time.Second)

	return scope
}

func NewMetricsHandler(ctx context.Context, p prom.Registerer, prefix string) client.MetricsHandler {
	return tallyhandler.NewMetricsHandler(NewPrometheusScope(ctx, p, prefix))
}
