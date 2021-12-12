package solr_search_client

//https://www.elastic.co/blog/how-to-add-powerful-search-existing-sql-applications-elasticsearch-video-tutorial
type ISolrearchClient interface {
}

type SolrSearchClient struct {
}

func New() *SolrSearchClient {

	return &SolrSearchClient{}
}
