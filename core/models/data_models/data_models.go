package data_models

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