package subscription_models

import (
	"time"

	"github.com/paulmsegeya/subscription/core/models/auth_models"
	"github.com/paulmsegeya/subscription/core/models/base_models"
	"github.com/paulmsegeya/subscription/core/models/workflow_models"
)



type Department struct {
	base_models.Foundation
	Code         string
	WorkflowList []*workflow_models.Workflow
}

type Subscription struct {
	base_models.Foundation
	SubscriberID uint
	StartDate    time.Time
	EndDate      time.Time
	CountDown    int // if count down =no of days in a year -less no of days in a month
	IsTrial      int8
	Price        float64
	Discount     float64
	Vat          float64
	ProductList  []*Product
	ServiceList  []*Service
	ScheduleList []*Schedule
	SupportList  []*Support
	WorkflowList []*workflow_models.Workflow
}

type Subscriber struct {
	auth_models.User
	SubscriptionList []*Subscription
	WorkflowList     []*workflow_models.Workflow
}

type Service struct {
	base_models.Foundation
	SubscriptionID uint
	LimitationList []*Limitation
	LicenceList    []*Licence
	WorkflowList   []*workflow_models.Workflow
}
type Customer struct {
	auth_models.User
	WorkflowList []*workflow_models.Workflow
}

type Contact struct {
	base_models.Foundation
	Mobile       string
	Address      string
	Website      string
	Email        string
	Fax          string
	ShipperID    uint
	WorkflowList []*workflow_models.Workflow
}

type Product struct {
	base_models.Foundation
	SubscriptionID uint
	LimitationList []*Limitation
	LicenceList    []*Licence
	Price          float64
	Quantity       int
	WorkflowList   []*workflow_models.Workflow
}

type Limitation struct {
	base_models.Foundation
	ProductID    uint
	ServiceID    uint
	WorkflowList []*workflow_models.Workflow
}

type Notification struct {
	base_models.Foundation
	ScheduleID   uint
	WorkflowList []*workflow_models.Workflow
}

type Licence struct {
	base_models.Foundation
	ProductID    uint
	ServiceID    uint
	WorkflowList []*workflow_models.Workflow
}

type Schedule struct {
	base_models.Foundation
	SubscriptionID uint
	Notification   *Notification
	WorkflowList   []*workflow_models.Workflow
}

type Contract struct {
	base_models.Foundation
	AgreementList []*Agreement
	WorkflowList  []*workflow_models.Workflow
}
type Report struct {
	base_models.Foundation
	Title          string
	RunDate        time.Time
	ReportTemplate string
	ReportFile     []byte
	ReportFormat   string
	WorkflowList   []*workflow_models.Workflow
}
type Agreement struct {
	base_models.Foundation
	ContractID   uint
	WorkflowList []*workflow_models.Workflow
}

type Support struct {
	base_models.Foundation
	SubscriptionID uint
	Team           []*Team
	WorkflowList   []*workflow_models.Workflow
}

type Team struct {
	base_models.Foundation
	SupportID    uint
	StaffList    []*Staff
	WorkflowList []*workflow_models.Workflow
}

type Staff struct {
	auth_models.User
	TeamID       uint
	WorkflowList []*workflow_models.Workflow
}
