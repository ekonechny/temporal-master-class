package main

import (
	"context"
	"log"
	"net/http"
	"os"

	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"golang.org/x/sync/errgroup"

	"temporal-master-class/internal/generated/temporal"
	"temporal-master-class/internal/services/assortment"
	"temporal-master-class/internal/services/payment"
	"temporal-master-class/internal/services/vendors"
	"temporal-master-class/internal/workflows/checkout"
	ch "temporal-master-class/internal/workflows/checkout/activities"
	"temporal-master-class/internal/workflows/customer"
	ca "temporal-master-class/internal/workflows/customer/activities"
	"temporal-master-class/internal/workflows/processing"
	"temporal-master-class/internal/workflows/processing/activities"
	tcl_client "temporal-master-class/pkg/tcl-client"
)

func main() {
	reg := prom.NewRegistry()
	app, err := temporal.NewCustomerCli(
		temporal.NewCustomerCliOptions().
			WithClient(func(ctx *cli.Context) (client.Client, error) {
				return tcl_client.New(
					context.Background(),
					tcl_client.Config{},
					reg,
				)
			}).WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
			w := worker.New(c, temporal.CustomerTaskQueue, worker.Options{})
			temporal.RegisterCustomerFlowWorkflow(w, customer.Register)
			temporal.RegisterCheckoutFlowWorkflow(w, checkout.Register)
			temporal.RegisterProcessingFlowWorkflow(w, processing.Register)
			temporal.RegisterCustomerActivities(w,
				ca.Register(
					assortment.New(),
				),
			)
			temporal.RegisterCheckoutActivities(w,
				ch.Register(
					assortment.New(),
					payment.New(),
				),
			)
			temporal.RegisterProcessingActivities(w,
				activities.Register(
					payment.New(),
					vendors.New(),
				),
			)
			return w, nil
		}),
	)
	if err != nil {
		log.Fatalf("error initializing example cli: %v", err)
	}

	g := errgroup.Group{}
	g.Go(func() error {
		http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
		return http.ListenAndServe(":8081", nil)
	})
	g.Go(func() error {
		return app.Run(os.Args)
	})
	log.Fatal(g.Wait())
}
