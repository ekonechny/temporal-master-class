package processing

import (
	"errors"
	"time"

	"go.temporal.io/sdk/workflow"

	"temporal-master-class/internal/generated/temporal"
	"temporal-master-class/internal/utils"
)

func Register(ctx workflow.Context, input *temporal.ProcessingFlowWorkflowInput) (temporal.ProcessingFlowWorkflow, error) {
	return &Workflow{
		order: &temporal.Order{
			Id:          input.Req.Id,
			Customer:    input.Req.Customer,
			Cart:        input.Req.Cart,
			PaymentType: input.Req.PaymentType,
			CreatedAt:   utils.TimeToTimestamp(workflow.Now(ctx)),
		},
	}, nil
}

type Workflow struct {
	order *temporal.Order
}

func (w *Workflow) GetOrder() (*temporal.Order, error) {
	return w.order, nil
}

func (w *Workflow) Execute(ctx workflow.Context) error {
	// Ожидаем подтверждения оплаты для онлайн платежа
	if w.order.PaymentType == temporal.PaymentType_ONLINE {
		payment, err := temporal.GetPayment(ctx, &temporal.PaymentStatusRequest{})
		if err != nil {
			return err
		}
		if payment.Status == temporal.PaymentStatus_PaymentStatusFailed {
			return err
		}
	}

	// Создаем заказ в магазине для сборки
	vendorOrderResponse, err := temporal.CreateVendorOrder(ctx, &temporal.CreateVendorOrderRequest{})
	if err != nil {
		return err
	}

	// Ожидаем пока вендор соберет заказ
	for {
		if err = workflow.Sleep(ctx, time.Minute); err != nil {
			return err
		}
		vendorOrder, err := temporal.GetVendorOrder(ctx, &temporal.VendorOrderRequest{
			Id: vendorOrderResponse.Id,
		})
		if err != nil {
			return err
		}
		if vendorOrder.Status == temporal.VendorOrderStatus_VendorOrderCancelled {
			return errors.New("vendor cancel order")
		}
		if vendorOrder.Status > temporal.VendorOrderStatus_VendorOrderInDelivery {
			break
		}
	}

	return workflow.Sleep(ctx, time.Minute*5)
}
