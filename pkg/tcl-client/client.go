package tcl_client

import (
	"context"

	"github.com/prometheus/client_golang/prometheus"
	"go.temporal.io/sdk/client"
)

func New(ctx context.Context, c Config, p prometheus.Registerer) (client.Client, error) {
	logger, err := NewLogger()
	if err != nil {
		return nil, err
	}
	opts := client.Options{
		Logger:    logger,
		HostPort:  c.Host,
		Namespace: c.Namespace,
	}

	if p != nil {
		opts.MetricsHandler = NewMetricsHandler(ctx, p, c.Prefix)
	}

	return client.NewLazyClient(opts)
}
