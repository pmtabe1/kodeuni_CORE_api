package workflow_models

import "github.com/paulmsegeya/subscription/core/models/error_models"

type WorkflowRepositoryResponse struct {
	Workflow                *Workflow
	WorkflowList            []*Workflow
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}