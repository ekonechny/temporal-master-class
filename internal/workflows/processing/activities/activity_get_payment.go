package activities

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"

	"temporal-master-class/internal/generated/temporal"
)

// GetPayment - поллинг активити для получения статуса платежа каждые 15 секунд, пока он не сменит статус на New
// Пример отсюда:
// https://github.com/temporalio/samples-go/blob/main/polling/frequent/activities.go#L22
func (a *Activities) GetPayment(ctx context.Context, req *temporal.PaymentStatusRequest) (*temporal.PaymentStatusResponse, error) {
	for {
		resp, err := a.paymentClient.GetPayment(ctx, req.Id)
		if err == nil && resp != nil && resp.Status != temporal.PaymentStatus_PaymentStatusNew {
			return resp, nil
		}
		activity.RecordHeartbeat(ctx)
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(time.Second * 15):
		}
	}
}
