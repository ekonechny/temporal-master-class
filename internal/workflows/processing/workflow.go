package processing

import (
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
		confirm: input.VendorOrderConfirm,
	}, nil
}

type Workflow struct {
	order   *temporal.Order
	confirm *temporal.VendorOrderConfirmSignal
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
			w.order.Status = temporal.OrderStatus_OrderStatusCancelled
			return workflow.ErrCanceled
		}
	}

	// Создаем заказ в магазине для сборки
	vendorOrderResponse, err := temporal.CreateVendorOrder(ctx, &temporal.CreateVendorOrderRequest{})
	if err != nil {
		return err
	}

	// Ожидаем пока вендор соберет заказ
	for {
		if w.order.Status == temporal.OrderStatus_OrderStatusCancelled {
			return workflow.ErrCanceled
		}
		if w.order.Status > temporal.OrderStatus_OrderStatusReady {
			break
		}
		timerFuture := workflow.NewTimer(ctx, time.Minute*15)

		s := workflow.NewSelector(ctx)
		w.confirm.Select(s, func(request *temporal.VendorOrderConfirmRequest) {
			w.order.Status = vendorOrderStatusMap[request.Status]
		})
		s.AddFuture(timerFuture, func(workflow.Future) {
			vendorOrder, err := temporal.GetVendorOrder(ctx, &temporal.VendorOrderRequest{
				Id: vendorOrderResponse.Id,
			})
			if err != nil {
				workflow.GetLogger(ctx).Error("failed to poll", "err", err.Error())
			}
			w.order.Status = vendorOrderStatusMap[vendorOrder.Status]
		})

		s.Select(ctx)
	}

	return nil
}

var vendorOrderStatusMap = map[temporal.VendorOrderStatus]temporal.OrderStatus{
	temporal.VendorOrderStatus_VendorOrderStatusNew:        temporal.OrderStatus_OrderStatusConfirmed,
	temporal.VendorOrderStatus_VendorOrderStatusConfirmed:  temporal.OrderStatus_OrderStatusVendorConfirmed,
	temporal.VendorOrderStatus_VendorOrderStatusPicking:    temporal.OrderStatus_OrderStatusPicking,
	temporal.VendorOrderStatus_VendorOrderStatusReady:      temporal.OrderStatus_OrderStatusReady,
	temporal.VendorOrderStatus_VendorOrderInStatusDelivery: temporal.OrderStatus_OrderStatusInDelivery,
	temporal.VendorOrderStatus_VendorOrderStatusCancelled:  temporal.OrderStatus_OrderStatusCancelled,
}
