package activities

import (
	"context"

	"temporal-master-class/internal/generated/temporal"
)

func (a *Activities) CreatePayment(ctx context.Context, req *temporal.CreatePaymentRequest) (*temporal.CreatePaymentResponse, error) {
	return a.paymentClient.CreatePayment(ctx, req)
}
