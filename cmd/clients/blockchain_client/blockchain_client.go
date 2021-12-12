package blockchain_client

import (
	"github.com/paulmsegeya/pos/cmd/clients/blockchain_client/bitcoin_client"
	"github.com/paulmsegeya/pos/cmd/clients/blockchain_client/etherium_client"
)

type IBlockchainClient interface {
}

type BlockchainClient struct {
	BitcoinClient  *bitcoin_client.BitcoinClient
	EtheriumClient *etherium_client.EtheriumClient
}

func New() *BlockchainClient {

	return &BlockchainClient{
		BitcoinClient:  bitcoin_client.New(),
		EtheriumClient: etherium_client.New(),
	}
}
