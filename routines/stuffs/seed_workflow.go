package main

import (
	"context"
	"time"

	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// ApplicationName is the task list name
const ApplicationName = "SeedGroup"

// This is registration process where you register all your workflows
// and activity function handlers.
func init() {
	workflow.Register(SeedWorkflow)
	activity.Register(seedActivity)
}

// SeedWorkflow workflow decider
func SeedWorkflow(ctx workflow.Context, name string) error {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("seed workflow started")

	var seedResult string
	err := workflow.ExecuteActivity(ctx, seedActivity, name).Get(ctx, &seedResult)
	if err != nil {
		logger.Error("Activity failed.", zap.Error(err))
		return err
	}

	logger.Info("Workflow completed.", zap.String("Result", seedResult))

	return nil
}

func seedActivity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("seed activity started")
	return "reply: " + name + "!", nil
}
