package __root

import (
	"errors"
	"strings"
	"time"

	"temporal-master-class/generated/temporal"

	"go.temporal.io/sdk/workflow"
)

func Register(ctx workflow.Context, input *temporal.CustomerFlowWorkflowInput) (temporal.CustomerFlowWorkflow, error) {
	return &Workflow{
		req:                 input.Req,
		deleteProfileSignal: input.DeleteProfile,
		deleteCartSignal:    input.DeleteCart,
		setAddressSignal:    input.SetAddress,
	}, nil
}

type Workflow struct {
	req                 *temporal.CustomerFlowRequest
	deleteProfileSignal *temporal.DeleteProfileSignal
	deleteCartSignal    *temporal.DeleteCartSignal
	setAddressSignal    *temporal.SetAddressSignal

	profile *temporal.Profile
	cart    *temporal.Cart
}

func (w *Workflow) Execute(ctx workflow.Context) error {
	// Создаем профиль
	w.profile = &temporal.Profile{
		Id:    getProfileIdFromWorkflow(ctx),
		Name:  w.req.GetName(),
		Phone: w.req.GetPhone(),
	}
	workflow.GetLogger(ctx).Info("new profile", "profile", w.profile)
	// Ожидаем 1 минуту пока пользователь введет адрес или отменяем его "жизненный цикл"
	var isCancelled bool
	sel := workflow.NewSelector(ctx)

	w.setAddressSignal.Select(sel, func(request *temporal.SetAddressRequest) {
		w.profile.Address = request.Address
	})
	sel.AddFuture(workflow.NewTimer(ctx, time.Minute), func(f workflow.Future) {
		isCancelled = true
	})
	sel.Select(ctx)
	if isCancelled {
		return workflow.ErrCanceled
	}

	var stop bool
	for !stop {
		sel := workflow.NewSelector(ctx)
		w.setAddressSignal.Select(sel, func(request *temporal.SetAddressRequest) {
			w.profile.Address = request.Address
		})
		w.deleteProfileSignal.Select(sel, func() {
			stop = true
		})
		w.deleteCartSignal.Select(sel, func() {
			w.cart = nil
		})
		sel.Select(ctx)
	}

	return nil
}

func (w *Workflow) GetProfile() (*temporal.Profile, error) {
	return w.profile, nil
}

func (w *Workflow) GetCart() (*temporal.Cart, error) {
	return w.cart, nil
}

func (w *Workflow) UpdateProfile(context workflow.Context, request *temporal.UpdateProfileRequest) (*temporal.Profile, error) {
	w.profile.Name = request.Name
	return w.profile, nil
}

func (w *Workflow) UpdateCart(context workflow.Context, request *temporal.UpdateCartRequest) (*temporal.Cart, error) {
	// TODO: здесь мы будем запускать процессинг корзины, но пока это мок
	w.cart = &temporal.Cart{
		Products: request.Products,
		Total:    calculateTotal(request.Products),
	}
	return w.cart, nil
}

func (w *Workflow) Checkout(context workflow.Context, request *temporal.CheckoutRequest) (*temporal.Order, error) {
	if w.cart == nil {
		return nil, errors.New("cart is empty")
	}
	// TODO: Здесь мы будем создавать заказ, но пока это мок
	return &temporal.Order{
		Id:          "BC2B6F6D-F598-4D66-AF34-AD32BF9C7945",
		Customer:    w.profile,
		Cart:        w.cart,
		PaymentType: request.PaymentType,
	}, nil
}

func calculateTotal(products []*temporal.Product) int32 {
	var total int32
	for i := range products {
		total += products[i].Qty * products[i].Price
	}
	return total
}

func getProfileIdFromWorkflow(ctx workflow.Context) string {
	return strings.Split(workflow.GetInfo(ctx).WorkflowExecution.ID, "/")[1]
}
