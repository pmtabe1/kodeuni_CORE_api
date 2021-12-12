package search_client

import (
	"github.com/paulmsegeya/pos/cmd/clients/search_client/elastic_search_client"
	"github.com/paulmsegeya/pos/cmd/clients/search_client/solr_search_client"
)

type ISearchClient interface {
}

type SearchClient struct {
	ElasticSearchClient *elastic_search_client.ElasticSearchClient
	SolrSearchClient    *solr_search_client.SolrSearchClient
}

func New() *SearchClient {

	return &SearchClient{
		ElasticSearchClient: elastic_search_client.New("", "", "", ""),
		SolrSearchClient:    solr_search_client.New(),
	}
}
