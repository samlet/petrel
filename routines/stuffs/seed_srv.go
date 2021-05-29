package main

import (
	"time"

	"github.com/pborman/uuid"
	"github.com/samlet/petrel/routines/common"
	"go.uber.org/cadence/client"
	"go.uber.org/cadence/worker"
)

// This needs to be done as part of a bootstrap step when the process starts.
// The workers are supposed to be long running.
func startWorkers(h *common.SampleHelper) {
	// Configure worker options.
	workerOptions := worker.Options{
		MetricsScope: h.Scope,
		Logger:       h.Logger,
	}
	h.StartWorkers(h.Config.DomainName, ApplicationName, workerOptions)
}

func startSeedWorkflow(h *common.SampleHelper) {
	workflowOptions := client.StartWorkflowOptions{
		ID:                              "Seed_" + uuid.New(),
		TaskList:                        ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}
	h.StartWorkflow(workflowOptions, SeedWorkflow, "Cadence")
}
