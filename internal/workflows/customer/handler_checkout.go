package customer

import (
	"errors"

	"go.temporal.io/sdk/workflow"

	"temporal-master-class/internal/generated/temporal"
)

// Checkout - создает заказ через дочернее workflow, подробнее:
// https://docs.temporal.io/encyclopedia/child-workflows
func (w *Workflow) Checkout(ctx workflow.Context, request *temporal.CheckoutRequest) (*temporal.Order, error) {
	if w.cart == nil {
		return nil, errors.New("cart is empty")
	}
	return temporal.CheckoutFlowChild(ctx, &temporal.CheckoutFlowRequest{
		PaymentType: request.PaymentType,
		Cart:        w.cart,
		Customer:    w.profile,
	})
}
