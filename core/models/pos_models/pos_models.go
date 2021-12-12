package pos_models

import (
	"time"

	"github.com/paulmsegeya/pos/core/models/base_models"
	"github.com/paulmsegeya/pos/core/models/auth_models"

)

type Secret struct {
	base_models.Foundation
	ClientToken  string
	SessionID    string
	SecretID     string
	ClientID     string
	RealmID      string
	Realm        string
	Timeout      *time.Time
	MaxRefresh   *time.Time
	AuthType     string
	Domain       string
	Secret       string
	SecretKey    string
	IdentityKey  string
	Validated    bool
	WorkflowList []*Workflow
}

type Channel struct {
	base_models.Foundation
	Reference         string
	ReferenceNumber   int
	Source            string
	Destination       string
	Amount            float64
	FromAccount       string
	FromAccountNumber int
	FromAccounId      int
	Msisdn            string
	Intent            string
	WorkflowList      []*Workflow
}
type Report struct {
	base_models.Foundation
	Title          string
	RunDate        time.Time
	ReportTemplate string
	ReportFile     []byte
	ReportFormat   string
	WorkflowList   []*Workflow
}
type Order struct {
	base_models.Foundation
	ShipperID        uint
	Payment          *Payment
	InvoiceList      []*Invoice
	OrderItemList    []*OrderItem
	ShipToAddress    string
	ShipToCity       string
	ShipToCountry    string
	ShipToPostalCode string
	ShipToName       string
	WorkflowList     []*Workflow
}

type OrderItem struct {
	base_models.Foundation
	OrderID uint
}

type Purchase struct {
	base_models.Foundation
	TaxRate          float64
	Amount           float32
	AmountWithVat    float64
	Vat              float64
	Discount         float64
	PurchaseItemList []*PurchaseItem
	WorkflowList     []*Workflow
}

type Wallet struct {
	base_models.Foundation
	Account       string
	AccountID     uint
	DependantID   uint
	MemberID      uint
	VirtualCardID uint
	Balance       float64
	Source        string
	Destination   string
}

type PurchaseItem struct {
	base_models.Foundation
	PurchaseID uint
	Cost       float64
	Price      float64
	Discount   float64
	Category   string
	AccountSet string
	TaxGroup   string
	Quantity   int
}

type Sell struct {
	base_models.Foundation
	Total         float64
	TotalWithVat  float64
	TotalVat      float64
	TotalDiscount float64
	SellItemList  []*SellItem
	WorkflowList  []*Workflow
}

type SellItem struct {
	SellID uint
	base_models.Foundation
	UnitPrice  float64
	Vat        float64
	Discount   float64
	Quantity   float64
	AccountSet string
	GLAccount  string
	Category   string
}

type Department struct {
	base_models.Foundation
	Manager      string
	EmployeeList []*Employee
}

type Employee struct {
	auth_models.User
	DepartmentID uint
	Salary       float64
	SN           string
	SSN          string
	JoinedOn     time.Time
	ResignedOn   time.Time
	LeftOn       time.Time
	ReportsTo    string
}

type Change struct {
	base_models.Foundation
	Amount       float64
	WorkflowList []*Workflow
}

type Product struct {
	base_models.FoundationWithDocumentBytes
	ProductItemList []*ProductItem
	WorkflowList    []*Workflow
	Sku             string
	SN              string
}

type ProductItem struct {
	base_models.Foundation
	Code       string
	UnitPrice  float64
	Vat        float64
	ProductID  uint
	SupplierID uint
}

type Till struct {
	base_models.Foundation
	TillNumber   int
	AsignedUser  *auth_models.User
	WorkflowList []*Workflow
}

type Register struct {
	base_models.Foundation
	RegisterName string
	Amount       float64
	AsignedUser  *auth_models.User
	WorkflowList []*Workflow
}

type Sale struct {
	base_models.Foundation
	SaleItemList *SaleItem
	WorkflowList []*Workflow
}

type SaleItem struct {
	base_models.Foundation
	SaleID uint
}

type Reference struct {
	base_models.Foundation
	ReferenceNumber  uint
	GlobalCounter    uint
	DailyCounter     uint
	Reference        string
	BankReference    string
	ChequeReference  string
	ChannelReference string
	ReconReference   string
}

type Catalogue struct {
	base_models.Foundation
	WorkflowList []*Workflow
}

type Customer struct {
	auth_models.User
	WorkflowList []*Workflow
}

type Contact struct {
	base_models.Foundation
	Mobile       string
	Address      string
	Website      string
	Email        string
	Fax          string
	ShipperID    uint
	WorkflowList []*Workflow
}

type Shipper struct {
	CompanyName string
	base_models.Foundation
	OrderList    []*Order
	Contact      *Contact
	ShipingID    uint
	WorkflowList []*Workflow
}

type Shipping struct {
	base_models.Foundation
	Company      string
	ShipperList  []*Shipper
	WorkflowList []*Workflow
}

type Invoice struct {
	base_models.Foundation
	OrderID         uint
	InvoiceItemList []*InvoiceItem
	WorkflowList    []*Workflow
}

type InvoiceItem struct {
	InvoiceID uint
	base_models.Foundation
	AccountSet string
	UnitPrice  float64
	Category   string
	GlAccount  string
}

type Supplier struct {
	base_models.Foundation
	ProductItemList []*ProductItem
	WorkflowList    []*Workflow
}

type Payment struct {
	base_models.Foundation
	OrderID       uint
	PaymentDate   time.Time
	Timestamp     int64
	Amount        float64
	Reference     string
	PaymentMethod []*PaymentMethod
	WorkflowList  []*Workflow
}

type PaymentMethod struct {
	base_models.Foundation
	PaymentID    uint
	WorkflowList []*Workflow
}

type Member struct {
	auth_models.User
	DependantList []*Dependant
	VirtualCard   *VirtualCard
	WorkflowList  []*Workflow
	Wallet        *Wallet
}

type Dependant struct {
	auth_models.User
	MemberID     uint
	VirtualCard  *VirtualCard
	WorkflowList []*Workflow
	Wallet       *Wallet
}

type Eod struct {
	base_models.Foundation
	EodDate         time.Time
	Balance         float64
	Report          []byte
	SnapshotBalance float64
	ChangeBalance   float64
	WorkflowList    []*Workflow
}

type Transaction struct {
	base_models.Foundation
	Payment
	WorkflowList []*Workflow
}

type VirtualCard struct {
	base_models.Foundation
	MemberID     uint
	DependantID  uint
	Balance      float64
	SpendingList []*Spending
	Wallet       *Wallet
}

type Spending struct {
	base_models.Foundation
	VirtualCardID uint
	Amount        float64
	Reason        string
}

type Workflow struct {
	base_models.Foundation
	SourceCode      string
	PurchaseID      uint
	SecretID        uint
	ChannelID       uint
	ReportID        uint
	PromotionID     uint
	UtilizationID   uint
	MarketingID     uint
	ConsumptionID   uint
	SupplierID      uint
	SaleID          uint
	SellID          uint
	OrderID         uint
	ShippingID      uint
	CatalogueID     uint
	EmployeeID      uint
	ProductID       uint
	InvoiceID       uint
	MemberID        uint
	PaymentMethodID uint
	ShipperID       uint
	ChangeID        uint
	CustomerID      uint
	TillID          uint
	TransactionID   uint
	UserID          uint
	RegisterID      uint
	InventoryID     uint
	DependantID     uint
	EodID           uint
	TaxID           uint
	VirtualCardID   uint
	StoreID         uint
	PaymentID       uint
	ContactID       uint
	ProfileID       uint
	DatalogID       uint
}

type Datalog struct {
	base_models.Foundation
	Payload      string
	WorkflowList []*Workflow
}

type Store struct {
	base_models.Foundation
	InventoryList []*Inventory
	Location      string
	WorkflowList  []*Workflow
}

type Inventory struct {
	base_models.Foundation
	StoreID      uint
	WorkflowList []*Workflow
}

type Tax struct {
	base_models.Foundation
	TaxItemList  []*TaxItem
	WorkflowList []*Workflow
}

type TaxItem struct {
	base_models.Foundation
	TaxID uint
}

type Receipt struct {
	base_models.Foundation
}

type Promotion struct {
	base_models.Foundation
	WorkflowList []*Workflow
}

type Marketing struct {
	base_models.Foundation
	WorkflowList []*Workflow
}

type Consumption struct {
	base_models.Foundation
	UtilizationList []*Utilization
	WorkflowList    []*Workflow
}

type Utilization struct {
	base_models.Foundation
	Amount        float64
	ConsumptionID uint
	User          *auth_models.User
	WorkflowList  []*Workflow
}

type AppStore struct {
	base_models.Foundation
	AppList []*App
}

type App struct {
	base_models.Foundation
	AppStoreID uint
}



