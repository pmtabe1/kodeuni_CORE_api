package airflow_client


type IAirflowClient interface {
	
}

type AirflowClient struct {
	
}

func New() *AirflowClient  {
	
	return &AirflowClient{}
}