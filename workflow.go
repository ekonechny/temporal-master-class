package temporal_master_class

import (
	"go.temporal.io/sdk/workflow"

	"temporal-master-class/generated/temporal"
)

func Register(_ workflow.Context, input *temporal.CreateOrderWorkflowInput) (temporal.CreateOrderWorkflow, error) {
	return &Workflow{
		order: &temporal.Order{
			Id:       input.Req.Id,
			Address:  input.Req.Address,
			Products: input.Req.Products,
			Total:    calculateTotal(input.Req.Products),
		},
		delete: input.Delete,
	}, nil
}

type Workflow struct {
	delete *temporal.DeleteSignal
	order  *temporal.Order
}

func (w *Workflow) Execute(ctx workflow.Context) error {
	// Ожидаем удаления
	w.delete.Receive(ctx)
	return workflow.ErrCanceled
}

func (w *Workflow) Read() (*temporal.Order, error) {
	return w.order, nil
}

func (w *Workflow) Update(_ workflow.Context, req *temporal.UpdateOrderRequest) (*temporal.Order, error) {
	w.order.Address = req.Address
	w.order.Products = req.Products
	w.order.Total = calculateTotal(req.Products)
	return w.order, nil
}

func calculateTotal(products []*temporal.Product) int32 {
	var r int32
	for _, p := range products {
		r += p.Price * p.Qty
	}
	return r
}
