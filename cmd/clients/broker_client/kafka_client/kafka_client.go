package kafka_client

type IKafkaClient interface {
}

type KafkaClient struct {
}

func New() *KafkaClient {

	return &KafkaClient{}
}
