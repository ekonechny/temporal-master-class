package processing

import (
	"errors"
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
		confirm: input.VendorOrderCallback,
		payment: workflow.NewChannel(ctx),
	}, nil
}

var (
	CustomerPhone   = t.NewSearchAttributeKeyString("CustomerPhone")
	CustomerId      = t.NewSearchAttributeKeyString("CustomerId")
	CustomerAddress = t.NewSearchAttributeKeyString("CustomerAddress")
)

type Workflow struct {
	order   *temporal.Order
	confirm *temporal.VendorOrderCallbackSignal
	payment workflow.Channel
}

func (w *Workflow) PaymentCallback(ctx workflow.Context, request *temporal.PaymentCallbackRequest) error {
	if w.order.Status > temporal.OrderStatus_OrderStatusNew {
		return errors.New("order status is not new")
	}
	w.payment.Send(ctx, request)
	return nil
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
		var (
			err           error
			paymentStatus temporal.PaymentStatus
		)

		// Запускаем футур с высокочастотным актитивити
		getPaymentCtx, getPaymentCancel := workflow.WithCancel(ctx)
		getPayment := temporal.GetPaymentAsync(getPaymentCtx, &temporal.PaymentStatusRequest{})
		getPaymentCallback := func(f workflow.Future) {
			var payment temporal.PaymentStatusResponse
			err = f.Get(ctx, &payment)
			if err != nil {
				return
			}
			paymentStatus = payment.Status
		}

		// Запускаем ресивер из updateHandler
		paymentReceiveCallback := func(f workflow.ReceiveChannel, _ bool) {
			var payment temporal.PaymentCallbackRequest
			_ = f.Receive(ctx, &payment)
			paymentStatus = payment.Status
		}

		// Запускаем селектор:
		// - Либо вернется результат активити, поллящий платежную систему
		// - Либо придет колбек от платежной системы
		workflow.
			NewSelector(ctx).
			AddFuture(getPayment.Future, getPaymentCallback).
			AddReceive(w.payment, paymentReceiveCallback).
			Select(ctx)

		// Дергаем cancelHandler для getPaymentActivity
		getPaymentCancel()

		// Если активити вернулось с ошибкой
		if err != nil {
			w.order.Status = temporal.OrderStatus_OrderStatusCancelled
			metrics.RecordOrderStatus(ctx, w.order)
			return workflow.ErrCanceled
		}

		// И если статус, что платеж отменен
		if paymentStatus == temporal.PaymentStatus_PaymentStatusFailed {
			w.order.Status = temporal.OrderStatus_OrderStatusCancelled
			metrics.RecordOrderStatus(ctx, w.order)
			return workflow.ErrCanceled
		}
	}

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

		// обработка колбека от вендора
		onVendorOrderCallback := func(request *temporal.VendorOrderCallbackRequest) {
			newStatus := vendorOrderStatusMap[request.Status]
			if w.order.Status == newStatus {
				return
			}
			w.order.Status = newStatus
			metrics.RecordOrderStatus(ctx, w.order)

		}

		// Если истекло время, то делаем сами запрос к вендору
		timerFuture := workflow.NewTimer(ctx, time.Minute)
		onTimeoutCallback := func(workflow.Future) {
			vendorOrder, err := temporal.GetVendorOrder(ctx, &temporal.VendorOrderRequest{
				Id: vendorOrderResponse.Id,
			})
			if err != nil {
				workflow.GetLogger(ctx).Error("failed to poll", "err", err.Error())
			}
			newStatus := vendorOrderStatusMap[vendorOrder.Status]
			if w.order.Status == newStatus {
				return
			}
			w.order.Status = newStatus
			metrics.RecordOrderStatus(ctx, w.order)

		}

		// Запускаем селектор:
		// - либо получаем колбек от вендора
		// - либо ждем миинуту минут
		w.confirm.
			Select(workflow.NewSelector(ctx), onVendorOrderCallback).
			AddFuture(timerFuture, onTimeoutCallback).
			Select(ctx)
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
