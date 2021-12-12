package flow_client

type IFlowClient interface {
}

type FlowClient struct {
}

func New() *FlowClient {

	return &FlowClient{}
}
