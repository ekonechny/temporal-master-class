package temporal_master_class

import (
	"strings"

	"go.temporal.io/sdk/workflow"

	"temporal-master-class/generated/temporal"
)

func Register(ctx workflow.Context, input *temporal.CreateWorkflowInput) (temporal.CreateWorkflow, error) {
	return &Workflow{
		profile: &temporal.Profile{
			Id:    getProfileIdFromWorkflow(ctx),
			Name:  input.Req.Name,
			Phone: input.Req.Phone,
		},
		delete: input.Delete,
	}, nil
}

type Workflow struct {
	delete  *temporal.DeleteSignal
	profile *temporal.Profile
}

func (w *Workflow) Execute(ctx workflow.Context) error {
	// Ожидаем удаления
	w.delete.Receive(ctx)
	return workflow.ErrCanceled
}

func (w *Workflow) Read() (*temporal.Profile, error) {
	return w.profile, nil
}

func (w *Workflow) Update(_ workflow.Context, request *temporal.UpdateRequest) (*temporal.Profile, error) {
	w.profile.Name = request.Name
	return w.profile, nil
}

func getProfileIdFromWorkflow(ctx workflow.Context) string {
	return strings.Split(workflow.GetInfo(ctx).WorkflowExecution.ID, "/")[1]
}
