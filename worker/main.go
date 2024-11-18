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
	// Инициализируем Worker сразу из кодгена с помощью cli
	app, err := temporal.NewOrderCli(
		temporal.NewOrderCliOptions().WithWorker(
			func(cmd *cli.Context, c client.Client) (worker.Worker, error) {
				w := worker.New(c, temporal.OrderTaskQueue, worker.Options{})
				temporal.RegisterCreateOrderWorkflow(w, tmc.Register)
				return w, nil
			}),
	)
	if err != nil {
		log.Fatalf("error initializing example cli: %v", err)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
