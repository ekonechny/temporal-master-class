package customer

import (
	"time"

	"go.temporal.io/sdk/workflow"

	"temporal-master-class/internal/generated/temporal"
)

func Register(ctx workflow.Context, input *temporal.CustomerFlowWorkflowInput) (temporal.CustomerFlowWorkflow, error) {
	return &Workflow{
		req:                 input.Req,
		deleteProfileSignal: input.DeleteProfile,
		deleteCartSignal:    input.DeleteCart,
		setAddressSignal:    input.SetAddress,
		profile: &temporal.Profile{
			//Id:    utils.WorkflowID(ctx),
			Name:  input.Req.GetName(),
			Phone: input.Req.GetPhone(),
		},
	}, nil
}

type Workflow struct {
	req                 *temporal.CustomerFlowRequest
	deleteProfileSignal *temporal.DeleteProfileSignal
	deleteCartSignal    *temporal.DeleteCartSignal
	setAddressSignal    *temporal.SetAddressSignal
	profile             *temporal.Profile
	cart                *temporal.Cart
}

func (w *Workflow) Execute(ctx workflow.Context) error {
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

func (w *Workflow) UpdateProfile(ctx workflow.Context, request *temporal.UpdateProfileRequest) (*temporal.Profile, error) {
	w.profile.Name = request.Name
	return w.profile, nil
}
