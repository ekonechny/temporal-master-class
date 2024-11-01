package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

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
)

func main() {
	app, err := temporal.NewCustomerCli(
		temporal.NewCustomerCliOptions().WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
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

	// run cli
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
