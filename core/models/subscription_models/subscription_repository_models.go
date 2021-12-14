package subscription_models

import "github.com/paulmsegeya/subscription/core/models/error_models"

type SubscriptionRepositoryResponse struct {
	Subscription            *Subscription
	SubscriptionList        []*Subscription
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type ContactRepositoryResponse struct {
	Contact                 *Contact
	ContactList             []*Contact
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type DepartmentRepositoryResponse struct {
	Department              *Department
	DepartmentList          []*Department
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type AgreementRepositoryResponse struct {
	Agreement               *Agreement
	AgreementList           []*Agreement
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type LimitationRepositoryResponse struct {
	Limitation              *Limitation
	LimitationList          []*Limitation
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type ContractRepositoryResponse struct {
	Contract                *Contract
	ContractList            []*Contract
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type NotificationRepositoryResponse struct {
	Notification            *Notification
	NotificationList        []*Notification
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type CustomerRepositoryResponse struct {
	Customer                *Customer
	CustomerList            []*Customer
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}
type ReportRepositoryResponse struct {
	Report                  *Report
	ReportList              []*Report
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type TeamRepositoryResponse struct {
	Team                    *Team
	TeamList                []*Team
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type SupportRepositoryResponse struct {
	Support                 *Support
	SupportList             []*Support
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type RepositoryResponse struct {
	Team                    *Team
	TeamList                []*Team
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type StaffRepositoryResponse struct {
	Staff                   *Staff
	StaffList               []*Staff
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type SubscriberRepositoryResponse struct {
	Subscriber              *Subscriber
	SubscriberList          []*Subscriber
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type ProductRepositoryResponse struct {
	Product                 *Product
	ProductList             []*Product
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type ServiceRepositoryResponse struct {
	Service                 *Service
	ServiceList             []*Service
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}
type LicenceRepositoryResponse struct {
	Licence                 *Licence
	LicenceList             []*Licence
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

// type LimitationRepositoryResponse struct {
// 	Licence                 *Licence
// 	LicenceList             []*Licence
// 	StatusCode              int
// 	RepositoryStatus        bool
// 	Error                   string
// 	Message                 string
// 	RepositoryErrorResponse *error_models.ErrorModel
// }

type ScheduleRepositoryResponse struct {
	Schedule                *Schedule
	ScheduleList            []*Schedule
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

