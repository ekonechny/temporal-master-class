package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"temporal-master-class/generated/temporal"

	tmc "temporal-master-class"
)

func main() {
	app, err := temporal.NewCustomerCli(
		temporal.NewCustomerCliOptions().WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
			w := worker.New(c, temporal.CustomerTaskQueue, worker.Options{})
			temporal.RegisterCustomerFlowWorkflow(w, tmc.Register)
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
