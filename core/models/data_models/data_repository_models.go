package data_models

import "github.com/paulmsegeya/subscription/core/models/error_models"

type DatalogRepositoryResponse struct {
	Datalog                 *Datalog
	DatalogList             []*Datalog
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}
