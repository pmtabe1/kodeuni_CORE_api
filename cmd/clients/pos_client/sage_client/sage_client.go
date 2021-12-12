package sage_client

import (
	"github.com/paulmsegeya/pos/cmd/clients/go_client"
	"github.com/paulmsegeya/pos/config/pos_config"
)

type IPosClient interface {
}

type PosClient struct {
	*go_client.GoClient
	*pos_config.PosConfig
}

func New() *PosClient {

	return &PosClient{
		GoClient:   go_client.New(),
		PosConfig: pos_config.New(),
	}
}
