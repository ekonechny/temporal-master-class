package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	tmc "temporal-master-class"
	"temporal-master-class/generated/temporal"
)

func main() {
	app, err := temporal.NewOrderCli(
		temporal.NewOrderCliOptions().WithWorker(func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
			w := worker.New(c, temporal.OrderTaskQueue, worker.Options{})
			//pb.RegisterOrderActivities(w, &crud.Activity{})
			temporal.RegisterCreateOrderWorkflow(w, tmc.Register)
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
