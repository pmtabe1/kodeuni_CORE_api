package sage_300_client

import "github.com/paulmsegeya/pos/cmd/clients/go_client"

type ISage300Client struct {
}

type Sage300Client struct {
	*go_client.GoClient
}

func New() *Sage300Client {

	return &Sage300Client{
		GoClient: go_client.New(),
	}
}
