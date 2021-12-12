package rpc_client

import "github.com/paulmsegeya/pos/cmd/clients/rpc_client/grpc_client"



type IRpcClient interface {
	
}

type RpcClient struct {
	
	GrpClient  *grpc_client.GrpcClient
}


func New()  *RpcClient{
	
	return &RpcClient{
		GrpClient: grpc_client.New(),
	}
}