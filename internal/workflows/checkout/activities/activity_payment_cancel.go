package activities

import (
	"context"

	"temporal-master-class/internal/generated/temporal"
)

func (a *Activities) PaymentCancel(ctx context.Context, req *temporal.PaymentCancelRequest) error {
	return a.paymentClient.PaymentCancel(ctx, req.Id)
}
