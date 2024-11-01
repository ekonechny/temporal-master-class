package temporal_master_class

import (
	"go.temporal.io/sdk/workflow"

	"temporal-master-class/generated/temporal"
)

func Register(ctx workflow.Context, input *temporal.CreateOrderWorkflowInput) (temporal.CreateOrderWorkflow, error) {
	return &Workflow{
		req:    input.Req,
		delete: input.Delete,
	}, nil
}

type Workflow struct {
	req    *temporal.CreateOrderRequest
	delete *temporal.DeleteSignal
	order  *temporal.Order
}

func (w *Workflow) Execute(ctx workflow.Context) error {
	w.order = &temporal.Order{
		Id:       w.req.GetId(),
		Address:  w.req.GetAddress(),
		Products: w.req.GetProducts(),
		Total:    calculateTotal(w.req.GetProducts()),
	}
	// Ожидаем удаления
	w.delete.Receive(ctx)
	return workflow.ErrCanceled
}

func (w *Workflow) Read() (*temporal.Order, error) {
	return w.order, nil
}

func (w *Workflow) Update(_ workflow.Context, req *temporal.UpdateOrderRequest) (*temporal.Order, error) {
	w.order.Address = req.GetAddress()
	w.order.Products = req.GetProducts()
	w.order.Total = calculateTotal(req.GetProducts())
	return w.order, nil
}

func calculateTotal(products []*temporal.Product) int32 {
	var r int32
	for _, p := range products {
		r += p.Price * p.Qty
	}
	return r
}
