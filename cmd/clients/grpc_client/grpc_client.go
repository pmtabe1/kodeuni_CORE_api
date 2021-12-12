package grpc_client

type GrpcClient struct {
}

type IGrpcClient interface {
}

func NewGrpcClient() *GrpcClient {
	return &GrpcClient{}
}
