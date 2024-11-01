package __helloworld

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

// Workflow is a HelloActivity World workflow definition.
func Workflow(ctx workflow.Context, name string) error {

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	})

	var helloResult string
	if err := workflow.ExecuteActivity(ctx, HelloActivity, name).Get(ctx, &helloResult); err != nil {
		return err
	}

	isReplay := workflow.IsReplaying(ctx)
	workflow.GetLogger(ctx).Info(fmt.Sprintf("isReplay %v", isReplay))

	var byeResult string
	if err := workflow.ExecuteActivity(ctx, ByeActivity, name).Get(ctx, &byeResult); err != nil {
		return err
	}

	//
	//
	//logger := workflow.GetLogger(ctx)
	//logger.Info("HelloWorld workflow started", "name", name)
	//
	//var helloResult string
	//if err := workflow.ExecuteActivity(ctx, HelloActivity, name).Get(ctx, &helloResult); err != nil {
	//	logger.Error("HelloActivity failed.", "Error", err)
	//	return err
	//}
	//
	//logger.Info("HelloActivity activity completed.", "result", helloResult)
	//
	//var byeResult string
	//if err := workflow.ExecuteActivity(ctx, ByeActivity, name).Get(ctx, &byeResult); err != nil {
	//	logger.Error("ByeActivity failed.", "Error", err)
	//	return err
	//}
	//
	//logger.Info("ByeActivity activity completed.", "result", byeResult)
	//
	//logger.Info("HelloWorkflow completed.")

	return nil
}

func HelloActivity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("HelloActivity", "name", name)
	return fmt.Sprintf("HelloActivity, %s!", name), nil
}

func ByeActivity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ByeActivity", "name", name)
	return fmt.Sprintf("ByeActivity, %s!", name), nil
}
