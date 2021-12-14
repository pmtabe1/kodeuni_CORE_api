package auth_models

import "github.com/paulmsegeya/subscription/core/models/error_models"

type SignupRepositoryResponse struct {
	Signup                  *Signup
	SignupList              []*Signup
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type SecretRepositoryResponse struct {
	Secret                  *Secret
	SecretList              []*Secret
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type LoginRepositoryResponse struct {
	Login                   *Login
	LoginList               []*Login
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type UserRepositoryResponse struct {
	User                    *User
	UserList                []*User
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type AclRepositoryResponse struct {
	Acl                     *Acl
	AclList                 []*Acl
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type PermissionRepositoryResponse struct {
	Permission              *Permission
	PermissionList          []*Permission
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type RoleRepositoryResponse struct {
	Role                    *Role
	RoleList                []*Role
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type GroupRepositoryResponse struct {
	Group                   *Group
	GroupList               []*Group
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}
