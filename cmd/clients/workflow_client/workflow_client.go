package workflow_client

import (
	"github.com/paulmsegeya/pos/cmd/clients/workflow_client/airflow_client"
	"github.com/paulmsegeya/pos/cmd/clients/workflow_client/camunda_client"
	"github.com/paulmsegeya/pos/cmd/clients/workflow_client/flow_client"
)

type IWorkflowClient interface {
}

type WorkflowClient struct {
	CamundaClient *camunda_client.CamundaClient
	AirflowClient *airflow_client.AirflowClient
	FlowClient    *flow_client.FlowClient
}

func New() *WorkflowClient {

	return &WorkflowClient{
		CamundaClient: camunda_client.New(),
		AirflowClient: airflow_client.New(),
		FlowClient:    flow_client.New(),
	}
}
