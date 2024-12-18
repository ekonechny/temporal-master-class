package utils

import (
	"strings"
	"time"

	"go.temporal.io/sdk/workflow"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func WorkflowID(ctx workflow.Context) string {
	if len(strings.Split(workflow.GetInfo(ctx).WorkflowExecution.ID, "/")) < 2 {
		return workflow.GetInfo(ctx).WorkflowExecution.ID
	}
	return strings.Split(workflow.GetInfo(ctx).WorkflowExecution.ID, "/")[1]
}

func TimeToTimestamp(t time.Time) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{
		Seconds: int64(t.Second()),
		Nanos:   int32(t.Nanosecond()),
	}
}
