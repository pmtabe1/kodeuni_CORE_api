package workflow_models

import "github.com/paulmsegeya/subscription/core/models/base_models"

type Workflow struct {
	base_models.Foundation
	SourceCode     string
	NotificationID uint
	AgreementID    uint
	SecretID       uint
	ReportID       uint
	TeamID         uint
	StaffID        uint
	SupportID      uint
	DepartmentID   uint
	SubscriberID   uint
	ContractID     uint
	ScheduleID     uint
	LimitationID   uint
	ContactID      uint
	LicenceID      uint
	SubscriptionID uint
	UserID         uint
	ProfileID      uint
	DatalogID      uint
	ServiceID      uint
	ProductID      uint
	CustomerID     uint
	GroupID        uint
	AclID          uint
	PermissionID   uint
	RoleID         uint
}
