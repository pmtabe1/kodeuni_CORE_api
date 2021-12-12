package elastic_search_client

import (

	// Import the Elasticsearch library packages
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	// "github.com/elastic/go-elasticsearch/v8/esapi"
)

type IElasticSearchClient interface {
}

type ElasticSearchClient struct {
	*elasticsearch.Client
}

func New(host string, port string, username string, password string) *ElasticSearchClient {

	var elasticSearchUri string

	if len(username) == 0 {
		username = "user"
	}

	if len(password) == 0 {
		password = "pass"
	}

	if len(host) == 0 {
		host = "http://localhost"
	}

	if len(port) == 0 {
		port = "9200"
	}

	//elasticSearchUriTemplate := fmt.Sprintf("%v:%v","%v","%v")

	// Declare an Elasticsearch configuration

	elasticSearchUri = fmt.Sprintf("%v:%v", host, port)

	cfg := elasticsearch.Config{
		Addresses: []string{
			elasticSearchUri,
		},
		Username: username,
		Password: password,
	}

	// Instantiate a new Elasticsearch client object instance
	client, err := elasticsearch.NewClient(cfg)

	if err != nil {
		fmt.Println("Elasticsearch connection error:", err)
	}

	return &ElasticSearchClient{
		Client: client,
	}
}
