package customer_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
	"go.temporal.io/sdk/workflow"

	"temporal-master-class/internal/generated/temporal"
	"temporal-master-class/internal/workflows/customer"
)

type UnitTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite

	env *testsuite.TestWorkflowEnvironment
}

func (s *UnitTestSuite) SetupTest() {
	s.env = s.NewTestWorkflowEnvironment()
	temporal.RegisterCustomerFlowWorkflow(s.env, customer.Register)
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}

func (s *UnitTestSuite) AfterTest(_, _ string) {
	s.env.AssertExpectations(s.T())
}

func (s *UnitTestSuite) Test_HappyPath() {
	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(temporal.SetAddressSignalName, &temporal.SetAddressRequest{Address: &temporal.Address{
			Title: "Санкт-Петербург, Невский проспект, 75",
			Lat:   "59.9342802",
			Long:  "30.3350986",
		}})
	}, time.Second)
	s.env.RegisterDelayedCallback(func() {
		s.env.SignalWorkflow(temporal.DeleteProfileSignalName, nil)
	}, time.Minute)
	s.env.ExecuteWorkflow(temporal.CustomerFlowWorkflowName, &temporal.CustomerFlowRequest{
		Name:  "Джон Персиков",
		Phone: "79650843621",
	}, workflow.RegisterOptions{})
	res, err := s.env.QueryWorkflow(temporal.GetProfileQueryName)
	s.NoError(err)
	var profile temporal.Profile
	err = res.Get(&profile)
	s.NoError(err)
	s.EqualValues("Джон Персиков", profile.Name)
	s.EqualValues("Санкт-Петербург, Невский проспект, 75", profile.Address.Title)
	s.True(s.env.IsWorkflowCompleted())
	s.NoError(s.env.GetWorkflowError())
}
