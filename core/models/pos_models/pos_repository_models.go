package pos_models

import "github.com/paulmsegeya/pos/core/models/error_models"

type ReferenceRepositoryResponse struct {
	Reference               *Reference
	ReferenceList           []*Reference
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type MarketingRepositoryResponse struct {
	Marketing               *Marketing
	MarketingList           []*Marketing
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type UtilizationRepositoryResponse struct {
	Utilization             *Utilization
	UtilizationList         []*Utilization
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type PromotionRepositoryResponse struct {
	Promotion               *Promotion
	PromotionList           []*Promotion
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type ConsumptionRepositoryResponse struct {
	Consumption             *Consumption
	ConsumptionList         []*Consumption
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

type ChannelRepositoryResponse struct {
	Channel                 *Channel
	ChannelList             []*Channel
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

type PaymentRepositoryResponse struct {
	Payment                 *Payment
	PaymentList             []*Payment
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type TransactionRepositoryResponse struct {
	Transaction             *Transaction
	TransactionList         []*Transaction
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type DatalogRepositoryResponse struct {
	Datalog                 *Datalog
	DatalogList             []*Datalog
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type CatalogueRepositoryResponse struct {
	Catalogue               *Catalogue
	CatalogueList           []*Catalogue
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type StoreRepositoryResponse struct {
	Store                   *Store
	StoreList               []*Store
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

type DependantRepositoryResponse struct {
	Dependant               *Dependant
	DependantList           []*Dependant
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type MemmberRepositoryResponse struct {
	Member                  *Member
	MemberList              []*Member
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type EodRepositoryResponse struct {
	Eod                     *Eod
	EodList                 []*Eod
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type InvoiceRepositoryResponse struct {
	Invoice                 *Invoice
	InvoiceList             []*Invoice
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}
type InvoiceItemRepositoryResponse struct {
	InvoiceItem             *InvoiceItem
	InvoiceItemList         []*InvoiceItem
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type MemberRepositoryResponse struct {
	Member                  *Member
	MemberList              []*Member
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type PurchaseRepositoryResponse struct {
	Purchase                *Purchase
	PurchaseList            []*Purchase
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}
type OrderRepositoryResponse struct {
	Order                   *Order
	OrderList               []*Order
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type OrderItemRepositoryResponse struct {
	OrderItem               *OrderItem
	OrderItemList           []*OrderItem
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type ProductItemRepositoryResponse struct {
	ProductItem             *ProductItem
	ProductItemList         []*ProductItem
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type PaymentMethodRepositoryResponse struct {
	PaymentMethod           *PaymentMethod
	PaymentMethodList       []*PaymentMethod
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type SaleRepositoryResponse struct {
	Sale                    *Sale
	SaleList                []*Sale
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type SaleItemRepositoryResponse struct {
	SaleItem                *SaleItem
	SaleItemList            []*SaleItem
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type ShipperRepositoryResponse struct {
	Shipper                 *Shipper
	ShipperList             []*Shipper
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type ShippingRepositoryResponse struct {
	Shipping                *Shipping
	ShippingList            []*Shipping
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type SupplierRepositoryResponse struct {
	Supplier                *Supplier
	SupplierList            []*Supplier
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}
type SellRepositoryResponse struct {
	Sell                    *Sell
	SellList                []*Sell
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type SellItemRepositoryResponse struct {
	SellItem                *SellItem
	SellItemList            []*SellItem
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type TaxRepositoryResponse struct {
	Tax                     *Tax
	TaxList                 []*Tax
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type TaxItemRepositoryResponse struct {
	TaxItem                 *TaxItem
	TaxItemList             []*TaxItem
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type TillRepositoryResponse struct {
	Till                    *Till
	TillList                []*Till
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type VirtualCardRepositoryResponse struct {
	VirtualCard             *VirtualCard
	VirtualCardList         []*VirtualCard
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type AppStoreRepositoryResponse struct {
	AppStore                *AppStore
	AppStoreList            []*AppStore
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

type ChangeRepositoryResponse struct {
	Change                  *Change
	ChangeList              []*Change
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type ReceiptRepositoryResponse struct {
	Receipt                 *Receipt
	ReceiptList             []*Receipt
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}
type RegisterRepositoryResponse struct {
	Register                *Register
	RegisterList            []*Register
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type PurchaseItemRepositoryResponse struct {
	PurchaseItem            *PurchaseItem
	PurchaseItemList        []*PurchaseItem
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

type WorkflowRepositoryResponse struct {
	Workflow                *Workflow
	WorkflowList            []*Workflow
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}

type WalletRepositoryResponse struct {
	Wallet                  *Wallet
	WalletList              []*Wallet
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}
type EmployeeRepositoryResponse struct {
	Employee                *Employee
	EmployeeList            []*Employee
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}
type InventoryRepositoryResponse struct {
	Inventory               *Inventory
	InventoryList           []*Inventory
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

type AppRepositoryResponse struct {
	App                     *App
	AppList                 []*App
	StatusCode              int
	RepositoryStatus        bool
	Error                   string
	Message                 string
	RepositoryErrorResponse *error_models.ErrorModel
}
