package processing

import (
	"time"

	t "go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	"temporal-master-class/internal/generated/temporal"
	"temporal-master-class/internal/utils"
	"temporal-master-class/internal/workflows/processing/metrics"
)

func Register(ctx workflow.Context, input *temporal.ProcessingFlowWorkflowInput) (temporal.ProcessingFlowWorkflow, error) {
	// Сохраняем индексы для поиска
	if err := workflow.UpsertTypedSearchAttributes(ctx, CustomerPhone.ValueSet(input.Req.Customer.Phone)); err != nil {
		return nil, err
	}
	if err := workflow.UpsertTypedSearchAttributes(ctx, CustomerId.ValueSet(input.Req.Customer.Id)); err != nil {
		return nil, err
	}
	if err := workflow.UpsertTypedSearchAttributes(ctx, CustomerAddress.ValueSet(input.Req.Customer.Address.Title)); err != nil {
		return nil, err
	}
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

var (
	CustomerPhone   = t.NewSearchAttributeKeyString("CustomerPhone")
	CustomerId      = t.NewSearchAttributeKeyString("CustomerId")
	CustomerAddress = t.NewSearchAttributeKeyString("CustomerAddress")
)

type Workflow struct {
	order   *temporal.Order
	confirm *temporal.VendorOrderConfirmSignal
}

func (w *Workflow) GetOrder() (*temporal.Order, error) {
	return w.order, nil
}

func (w *Workflow) Execute(ctx workflow.Context) error {
	// Пишем метрику для нового заказа
	metrics.RecordOrderStatus(ctx, w.order)

	// Ожидаем подтверждения оплаты для онлайн платежа
	// Используем высокочастотное активити
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

	// Выставляем статус и пишем метрику
	w.order.Status = temporal.OrderStatus_OrderStatusConfirmed
	metrics.RecordOrderStatus(ctx, w.order)

	// Создаем заказ в магазине для сборки
	vendorOrderResponse, err := temporal.CreateVendorOrder(ctx, &temporal.CreateVendorOrderRequest{})
	if err != nil {
		return err
	}

	// Ожидаем пока вендор соберет заказ через низкочастотное активити
	// Другой способ реализации
	// https://github.com/temporalio/samples-go/blob/main/polling/infrequent/workflow.go#L26
	for {
		if w.order.Status == temporal.OrderStatus_OrderStatusCancelled {
			return workflow.ErrCanceled
		}
		if w.order.Status > temporal.OrderStatus_OrderStatusReady {
			break
		}
		timerFuture := workflow.NewTimer(ctx, time.Minute)
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
