package data_models

import (
	"github.com/paulmsegeya/subscription/core/models/base_models"
	"github.com/paulmsegeya/subscription/core/models/workflow_models"
)

type EndpointParams struct {
	Protocol         string
	Secured          bool
	DatabaseName     string
	HostOrIP         string
	BaseEndPoint     string
	Version          string
	Port             int
	ParamMap         map[string]interface{}
	TargetModule     string
	ResourceEndpoint string
}

type Datalog struct {
	base_models.Foundation
	Payload      string
	WorkflowList []*workflow_models.Workflow
}
