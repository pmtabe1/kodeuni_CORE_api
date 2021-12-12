package broker_client

import (
	"github.com/paulmsegeya/pos/cmd/clients/broker_client/kafka_client"
	"github.com/paulmsegeya/pos/cmd/clients/broker_client/mttq_client"
)

type IBrokerClient interface {
}

type BrokerClient struct {
	KafkaClient *kafka_client.KafkaClient
	MttqClient  *mttq_client.MttqClient
}

func New() *BrokerClient {

	return &BrokerClient{
		KafkaClient: kafka_client.New(),
		MttqClient:  mttq_client.New(),
	}
}
