package main

import (
	"context"
	"errors"
	"github.com/pborman/uuid"
	"github.com/samlet/petrel/routines/common"
	"github.com/samlet/petrel/routines/recovery/cache"
	"go.uber.org/cadence"
	"go.uber.org/cadence/.gen/go/shared"
	"go.uber.org/cadence/activity"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
	"time"
)

type (
	// Params is the input parameters to RecoveryWorkflow
	Params struct {
		ID          string
		Type        string
		Concurrency int
	}

	// ListOpenExecutionsResult is the result returned from listOpenExecutions activity
	ListOpenExecutionsResult struct {
		ID     string
		Count  int
		HostID string
	}

	// RestartParams are parameters extracted from StartWorkflowExecution history event
	RestartParams struct {
		Options client.StartWorkflowOptions
		State   UserState
	}

	// SignalParams are the parameters extracted from SignalWorkflowExecution history event
	SignalParams struct {
		Name string
		Data TripEvent
	}
)

// ClientKey is the key for lookup
type ClientKey int

const (
	// DomainName used for this sample
	DomainName = "samples-domain"

	// CadenceClientKey for retrieving cadence client from context
	CadenceClientKey ClientKey = iota
	// WorkflowExecutionCacheKey for retrieving executions cache from context
	WorkflowExecutionCacheKey
)

// HostID - Use a new uuid just for demo so we can run 2 host specific activity workers on same machine.
// In real world case, you would use a hostname or ip address as HostID.
var HostID = uuid.New()

var (
	// ErrCadenceClientNotFound when cadence client is not found on context
	ErrCadenceClientNotFound = errors.New("failed to retrieve cadence client from context")
	// ErrExecutionCacheNotFound when executions cache is not found on context
	ErrExecutionCacheNotFound = errors.New("failed to retrieve cache from context")
)

// This is registration process where you register all your workflows
// and activity function handlers.
func init() {
	workflow.RegisterWithOptions(RecoverWorkflow, workflow.RegisterOptions{Name: "RecoverWorkflow"})
	activity.Register(listOpenExecutions)
	activity.Register(recoverExecutions)
}

// RecoverWorkflow is the workflow implementation to recover TripWorkflow executions
func RecoverWorkflow(ctx workflow.Context, params Params) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Recover workflow started.")

	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: 10 * time.Minute,
		StartToCloseTimeout:    10 * time.Minute,
		HeartbeatTimeout:       time.Second * 30,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result ListOpenExecutionsResult
	err := workflow.ExecuteActivity(ctx, listOpenExecutions, params.Type).Get(ctx, &result)
	if err != nil {
		logger.Error("Failed to list open workflow executions.", zap.Error(err))
		return err
	}

	concurrency := 1
	if params.Concurrency > 0 {
		concurrency = params.Concurrency
	}

	if result.Count < concurrency {
		concurrency = result.Count
	}

	batchSize := result.Count / concurrency
	if result.Count%concurrency != 0 {
		batchSize++
	}

	// Setup retry policy for recovery activity
	info := workflow.GetInfo(ctx)
	expiration := time.Duration(info.ExecutionStartToCloseTimeoutSeconds) * time.Second
	retryPolicy := &cadence.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2,
		MaximumInterval:    10 * time.Second,
		ExpirationInterval: expiration,
		MaximumAttempts:    100,
	}
	ao = workflow.ActivityOptions{
		ScheduleToStartTimeout: expiration,
		StartToCloseTimeout:    expiration,
		HeartbeatTimeout:       time.Second * 30,
		RetryPolicy:            retryPolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	doneCh := workflow.NewChannel(ctx)
	for i := 0; i < concurrency; i++ {
		startIndex := i * batchSize

		workflow.Go(ctx, func(ctx workflow.Context) {
			err = workflow.ExecuteActivity(ctx, recoverExecutions, result.ID, startIndex, batchSize).Get(ctx, nil)
			if err != nil {
				logger.Error("Recover executions failed.", zap.Int("StartIndex", startIndex), zap.Error(err))
			} else {
				logger.Info("Recover executions completed.", zap.Int("StartIndex", startIndex))
			}

			doneCh.Send(ctx, "done")
		})
	}

	for i := 0; i < concurrency; i++ {
		doneCh.Receive(ctx, nil)
	}

	logger.Info("Workflow completed.", zap.Int("Result", result.Count))

	return nil
}

func listOpenExecutions(ctx context.Context, workflowType string) (*ListOpenExecutionsResult, error) {
	key := uuid.New()
	logger := activity.GetLogger(ctx)
	logger.Info("List all open executions of type.",
		zap.String("WorkflowType", workflowType),
		zap.String("HostID", HostID))

	cadenceClient, err := getCadenceClientFromContext(ctx)
	if err != nil {
		return nil, err
	}

	executionsCache := ctx.Value(WorkflowExecutionCacheKey).(cache.Cache)
	if executionsCache == nil {
		logger.Error("Could not retrieve cache from context.")
		return nil, ErrExecutionCacheNotFound
	}

	openExecutions, err := getAllExecutionsOfType(ctx, cadenceClient, workflowType)
	if err != nil {
		return nil, err
	}

	executionsCache.Put(key, openExecutions)
	return &ListOpenExecutionsResult{
		ID:     key,
		Count:  len(openExecutions),
		HostID: HostID,
	}, nil
}

func recoverExecutions(ctx context.Context, key string, startIndex, batchSize int) error {
	logger := activity.GetLogger(ctx)
	logger.Info("Starting execution recovery.",
		zap.String("HostID", HostID),
		zap.String("Key", key),
		zap.Int("StartIndex", startIndex),
		zap.Int("BatchSize", batchSize))

	executionsCache := ctx.Value(WorkflowExecutionCacheKey).(cache.Cache)
	if executionsCache == nil {
		logger.Error("Could not retrieve cache from context.")
		return ErrExecutionCacheNotFound
	}

	openExecutions := executionsCache.Get(key).([]*shared.WorkflowExecution)
	endIndex := startIndex + batchSize

	// Check if this activity has previous heartbeat to retrieve progress from it
	if activity.HasHeartbeatDetails(ctx) {
		var finishedIndex int
		if err := activity.GetHeartbeatDetails(ctx, &finishedIndex); err == nil {
			// we have finished progress
			startIndex = finishedIndex + 1 // start from next one.
		}
	}

	for index := startIndex; index < endIndex && index < len(openExecutions); index++ {
		execution := openExecutions[index]
		if err := recoverSingleExecution(ctx, execution.GetWorkflowId()); err != nil {
			logger.Error("Failed to recover execution.",
				zap.String("WorkflowID", execution.GetWorkflowId()),
				zap.Error(err))
			return err
		}

		// Record a heartbeat after each recovery of execution
		activity.RecordHeartbeat(ctx, index)
	}

	return nil
}

func recoverSingleExecution(ctx context.Context, workflowID string) error {
	logger := activity.GetLogger(ctx)
	cadenceClient, err := getCadenceClientFromContext(ctx)
	if err != nil {
		return err
	}

	execution := &shared.WorkflowExecution{
		WorkflowId: common.StringPtr(workflowID),
	}
	history, err := getHistory(ctx, execution)
	if err != nil {
		return err
	}

	if history == nil || len(history) == 0 {
		// Nothing to recover
		return nil
	}

	firstEvent := history[0]
	lastEvent := history[len(history)-1]

	// Extract information from StartWorkflowExecution parameters so we can start a new run
	params, err := extractStateFromEvent(workflowID, firstEvent)
	if err != nil {
		return err
	}

	// Parse the entire history and extract all signals so they can be replayed back to new run
	signals, err := extractSignals(history)
	if err != nil {
		return err
	}

	// First terminate existing run if already running
	if !isExecutionCompleted(lastEvent) {
		err := cadenceClient.TerminateWorkflow(ctx, execution.GetWorkflowId(), execution.GetRunId(), "Recover", nil)
		if err != nil {
			return err
		}
	}

	// Start new execution run
	newRun, err := cadenceClient.StartWorkflow(ctx, params.Options, "TripWorkflow", params.State)
	if err != nil {
		return err
	}

	// re-inject all signals to new run
	for _, s := range signals {
		cadenceClient.SignalWorkflow(ctx, execution.GetWorkflowId(), newRun.RunID, s.Name, s.Data)
	}

	logger.Info("Successfully restarted workflow.",
		zap.String("WorkflowID", execution.GetWorkflowId()),
		zap.String("NewRunID", newRun.RunID))

	return nil
}

func extractStateFromEvent(workflowID string, event *shared.HistoryEvent) (*RestartParams, error) {
	switch event.GetEventType() {
	case shared.EventTypeWorkflowExecutionStarted:
		attr := event.WorkflowExecutionStartedEventAttributes
		state, err := deserializeUserState(attr.Input)
		if err != nil {
			// Corrupted Workflow Execution State
			return nil, err
		}
		return &RestartParams{
			Options: client.StartWorkflowOptions{
				ID:                              workflowID,
				TaskList:                        attr.TaskList.GetName(),
				ExecutionStartToCloseTimeout:    time.Second * time.Duration(attr.GetExecutionStartToCloseTimeoutSeconds()),
				DecisionTaskStartToCloseTimeout: time.Second * time.Duration(attr.GetTaskStartToCloseTimeoutSeconds()),
				WorkflowIDReusePolicy:           client.WorkflowIDReusePolicyAllowDuplicate,
				//RetryPolicy: attr.RetryPolicy,
			},
			State: state,
		}, nil
	default:
		return nil, errors.New("Unknown event type")
	}
}

func extractSignals(events []*shared.HistoryEvent) ([]*SignalParams, error) {
	var signals []*SignalParams
	for _, event := range events {
		if event.GetEventType() == shared.EventTypeWorkflowExecutionSignaled {
			attr := event.WorkflowExecutionSignaledEventAttributes
			if attr.GetSignalName() == TripSignalName && attr.Input != nil && len(attr.Input) > 0 {
				signalData, err := deserializeTripEvent(attr.Input)
				if err != nil {
					// Corrupted Signal Payload
					return nil, err
				}

				signal := &SignalParams{
					Name: attr.GetSignalName(),
					Data: signalData,
				}
				signals = append(signals, signal)
			}
		}
	}

	return signals, nil
}

func isExecutionCompleted(event *shared.HistoryEvent) bool {
	switch event.GetEventType() {
	case shared.EventTypeWorkflowExecutionCompleted, shared.EventTypeWorkflowExecutionTerminated,
		shared.EventTypeWorkflowExecutionCanceled, shared.EventTypeWorkflowExecutionFailed,
		shared.EventTypeWorkflowExecutionTimedOut:
		return true
	default:
		return false
	}
}

func getAllExecutionsOfType(ctx context.Context, cadenceClient client.Client,
	workflowType string) ([]*shared.WorkflowExecution, error) {
	var openExecutions []*shared.WorkflowExecution
	var nextPageToken []byte
	for hasMore := true; hasMore; hasMore = len(nextPageToken) > 0 {
		resp, err := cadenceClient.ListOpenWorkflow(ctx, &shared.ListOpenWorkflowExecutionsRequest{
			Domain:          common.StringPtr(DomainName),
			MaximumPageSize: common.Int32Ptr(10),
			NextPageToken:   nextPageToken,
			StartTimeFilter: &shared.StartTimeFilter{
				EarliestTime: common.Int64Ptr(0),
				LatestTime:   common.Int64Ptr(time.Now().UnixNano()),
			},
			TypeFilter: &shared.WorkflowTypeFilter{
				Name: common.StringPtr(workflowType),
			},
		})
		if err != nil {
			return nil, err
		}

		for _, r := range resp.Executions {
			openExecutions = append(openExecutions, r.Execution)
		}

		nextPageToken = resp.NextPageToken
		activity.RecordHeartbeat(ctx, nextPageToken)
	}

	return openExecutions, nil
}

func getHistory(ctx context.Context, execution *shared.WorkflowExecution) ([]*shared.HistoryEvent, error) {
	cadenceClient, err := getCadenceClientFromContext(ctx)
	if err != nil {
		return nil, err
	}

	iter := cadenceClient.GetWorkflowHistory(ctx, execution.GetWorkflowId(), execution.GetRunId(), false,
		shared.HistoryEventFilterTypeAllEvent)
	var events []*shared.HistoryEvent
	for iter.HasNext() {
		event, err := iter.Next()
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func getCadenceClientFromContext(ctx context.Context) (client.Client, error) {
	logger := activity.GetLogger(ctx)
	cadenceClient := ctx.Value(CadenceClientKey).(client.Client)
	if cadenceClient == nil {
		logger.Error("Could not retrieve cadence client from context.")
		return nil, ErrCadenceClientNotFound
	}

	return cadenceClient, nil
}
