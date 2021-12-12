package telemetry_client

type ITelemetryClient interface {
}

type TelemetryClient struct {
}

func New() *TelemetryClient {

	return &TelemetryClient{}
}
