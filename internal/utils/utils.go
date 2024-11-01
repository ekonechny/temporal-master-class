package utils

import (
	"strings"

	"go.temporal.io/sdk/workflow"
)

func WorkflowID(ctx workflow.Context) string {
	parts := strings.Split(workflow.GetInfo(ctx).WorkflowExecution.ID, "/")
	if len(parts) < 2 {
		return workflow.GetInfo(ctx).WorkflowExecution.ID
	}
	return parts[1]
}
