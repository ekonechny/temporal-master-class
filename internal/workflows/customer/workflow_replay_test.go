package customer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/worker"

	"temporal-master-class/internal/generated/temporal"
	"temporal-master-class/internal/workflows/customer"
)

type ReplayTestSuite struct {
	suite.Suite
}

func Test_Replays(t *testing.T) {
	s := new(ReplayTestSuite)
	suite.Run(t, s)
}

func (s *ReplayTestSuite) SetupTest() {
}

func (s *ReplayTestSuite) Test_Replays() {
	tests := []struct {
		name   string
		replay string
	}{
		{
			"Happy path",
			"replaytests/happy_path.json",
		},
		{
			"With Checkout",
			"replaytests/happy_path_with_checkout.json",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			replayer := worker.NewWorkflowReplayer()
			temporal.RegisterCustomerFlowWorkflow(replayer, customer.Register)
			err := replayer.ReplayWorkflowHistoryFromJSONFile(nil, tt.replay)
			require.NoError(s.T(), err)
		})
	}
}
