package main

import (
	"context"
	"time"

	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

type (
	flowInfo struct {
		Owner    string
		Bindings map[string]interface{}
	}
)

// ApplicationName is the task list for this workflow
const ApplicationName = "DummyGroup"

/**
* Activities used by file processing Dummy workflow.
 */
const (
	downloadFileActivityName = "downloadFileActivity"
	processFileActivityName  = "processFileActivity"
)

// This is registration process where you register all your workflows
// and activity function handlers.
func init() {
	workflow.Register(DummyWorkflow)
	activity.RegisterWithOptions(
		downloadFileActivity,
		activity.RegisterOptions{Name: downloadFileActivityName},
	)
	activity.RegisterWithOptions(
		processFileActivity,
		activity.RegisterOptions{Name: processFileActivityName},
	)
}

// DummyWorkflow Workflow Decider.
func DummyWorkflow(ctx workflow.Context, fInfo *flowInfo) error {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	var err error

	// Run "downloadFile"
	err = workflow.ExecuteActivity(ctx, downloadFileActivityName, *fInfo).Get(ctx, &fInfo)
	if err != nil {
		logger.Error("Activity 'downloadFile' failed.", zap.Error(err))
		return err
	}

	// Run "processFile"
	err = workflow.ExecuteActivity(ctx, processFileActivityName, *fInfo).Get(ctx, &fInfo)
	if err != nil {
		logger.Error("Activity 'processFile' failed.", zap.Error(err))
		return err
	}

	logger.Info("Workflow completed.", zap.String("Result", "OK"))
	return nil
}

// downloadFile Activity.
func downloadFileActivity(ctx context.Context, fInfo flowInfo) (*flowInfo, error) {
	processedInfo := &flowInfo{Owner: fInfo.Owner, Bindings: fInfo.Bindings}
	return processedInfo, nil
}

// processFile Activity.
func processFileActivity(ctx context.Context, fInfo flowInfo) (*flowInfo, error) {
	processedInfo := &flowInfo{Owner: fInfo.Owner, Bindings: fInfo.Bindings}
	return processedInfo, nil
}
