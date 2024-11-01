package metrics

import (
	"time"

	"github.com/uber-go/tally/v4"
	tallySDK "go.temporal.io/sdk/contrib/tally"
	"go.temporal.io/sdk/workflow"

	"temporal-master-class/internal/generated/temporal"
)

const metricOrderStatus = "order_status"

var orderStatusDurationBuckets = tally.DurationBuckets{
	time.Minute * 5,
	time.Minute * 15,
	time.Minute * 30,
	time.Hour,
	time.Hour * 2,
	time.Hour * 4,
	time.Hour * 8,
	time.Hour * 24,
}

func RecordOrderStatus(ctx workflow.Context, order *temporal.Order) {
	if workflow.IsReplaying(ctx) {
		return
	}
	tallySDK.ScopeFromHandler(workflow.GetMetricsHandler(ctx)).
		Tagged(tagOrderStatus(order)).
		Histogram(metricOrderStatus, orderStatusDurationBuckets).
		RecordDuration(orderStatusDuration(ctx, order))
}

func orderStatusDuration(ctx workflow.Context, order *temporal.Order) time.Duration {
	return order.CreatedAt.AsTime().Sub(workflow.Now(ctx))
}

func tagOrderStatus(order *temporal.Order) map[string]string {
	return map[string]string{
		"status": order.Status.String(),
	}
}
