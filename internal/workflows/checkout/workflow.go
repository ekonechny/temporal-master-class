package checkout

import (
	"go.temporal.io/sdk/workflow"

	"temporal-master-class/internal/generated/temporal"
	"temporal-master-class/internal/utils"
)

func Register(ctx workflow.Context, input *temporal.CheckoutFlowWorkflowInput) (temporal.CheckoutFlowWorkflow, error) {
	return &Workflow{
		req: input.Req,
	}, nil
}

type Workflow struct {
	req *temporal.CheckoutFlowRequest
}

func (w *Workflow) Execute(ctx workflow.Context) (*temporal.Order, error) {
	// Резервируем продукты
	reserveProducts := make([]*temporal.AssortmentReserveProductRequest, 0, len(w.req.Cart.Products))
	for _, p := range w.req.Cart.Products {
		reserveProducts = append(reserveProducts, &temporal.AssortmentReserveProductRequest{
			Id:  p.Id,
			Qty: p.Qty,
		})
	}
	err := temporal.AssortmentReserve(ctx, &temporal.AssortmentReserveRequest{Products: reserveProducts})
	if err != nil {
		return nil, err
	}
	if w.req.PaymentType == temporal.PaymentType_ONLINE {
		_, err = temporal.CreatePayment(ctx, &temporal.CreatePaymentRequest{})
		if err != nil {
			return nil, err
		}
	}
	order := &temporal.Order{
		Id:          utils.WorkflowID(ctx),
		Customer:    w.req.Customer,
		Cart:        w.req.Cart,
		PaymentType: w.req.PaymentType,
	}

	// Здесь мы запускаем дочернее Workflow, но уже с политикой Abandon
	// https://docs.temporal.io/encyclopedia/child-workflows#parent-close-policy
	run, err := temporal.ProcessingFlowChildAsync(ctx, &temporal.ProcessingFlowRequest{
		Id:          order.Id,
		Customer:    order.Customer,
		Cart:        order.Cart,
		PaymentType: order.PaymentType,
	})
	if err != nil {
		return nil, err
	}
	_, err = run.WaitStart(ctx)
	if err != nil {
		return nil, err
	}

	return order, nil
}
