package main

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/encoded"
	"go.uber.org/cadence/testsuite"
)

type UnitTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}

func (s *UnitTestSuite) Test_DummyWorkflow() {
	owner := "test-id"
	expectedCall := []string{
		"downloadFileActivity",
		"processFileActivity",
	}

	var activityCalled []string
	env := s.NewTestWorkflowEnvironment()
	env.SetOnActivityStartedListener(func(activityInfo *activity.Info, ctx context.Context, args encoded.Values) {
		activityType := activityInfo.ActivityType.Name
		if strings.HasPrefix(activityType, "internalSession") {
			return
		}
		activityCalled = append(activityCalled, activityType)
		switch activityType {

		case expectedCall[0]:
			var input flowInfo
			s.NoError(args.Get(&input))
			s.Equal(input.Owner, owner)
		case expectedCall[1]:
			var input flowInfo
			s.NoError(args.Get(&input))
			s.Equal(input.Owner, owner)
		default:
			panic("unexpected activity call")
			// println("unexpected activity call -> "+activityType)
		}
	})

	env.ExecuteWorkflow(DummyWorkflow, &flowInfo{Owner: owner})

	s.True(env.IsWorkflowCompleted())
	s.NoError(env.GetWorkflowError())
	s.Equal(expectedCall, activityCalled)
}
