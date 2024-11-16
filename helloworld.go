package __helloworld

import (
	"context"
	"fmt"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

// Workflow - самое простейший рабочий процесс исполняющий два действия Hello и Bye
// калька отсюда https://github.com/temporalio/samples-go/blob/main/helloworld/helloworld.go
func Workflow(ctx workflow.Context, name string) error {

	// Создаем логер
	logger := workflow.GetLogger(ctx)
	logger.Info("HelloWorld workflow started", "name", name)

	// Выполяем HelloActivity
	var helloResult string
	if err := workflow.ExecuteActivity(ctx, HelloActivity, name).Get(ctx, &helloResult); err != nil {
		logger.Error("HelloActivity failed.", "Error", err)
		return err
	}
	logger.Info("HelloActivity activity completed.", "result", helloResult)

	// Выполняем ByeActivity
	var byeResult string
	if err := workflow.ExecuteActivity(ctx, ByeActivity, name).Get(ctx, &byeResult); err != nil {
		logger.Error("ByeActivity failed.", "Error", err)
		return err
	}
	logger.Info("ByeActivity activity completed.", "result", byeResult)

	// Завершаем исполнение
	logger.Info("HelloWorkflow completed.")
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
