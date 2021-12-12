package handlers

import (
	"github.com/paulmsegeya/pos/internal/handlers/acl_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/app_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/appstore_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/auth_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/catalogue_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/change_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/channel_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/consumption_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/contact_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/customer_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/datalog_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/department_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/dependant_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/employee_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/eod_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/group_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/inventory_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/invoice_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/marketing_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/member_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/order_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/payment_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/payment_method_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/permission_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/product_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/promotion_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/purchase_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/reference_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/register_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/report_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/role_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/sale_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/sell_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/shipper_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/shipping_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/store_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/supplier_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/tax_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/till_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/transaction_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/user_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/utilization_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/virtualcard_handlers"
	"github.com/paulmsegeya/pos/internal/handlers/wallet_handlers"

	"github.com/paulmsegeya/pos/internal/handlers/workflow_handlers"
)

type IHandlers interface {
}

type Handlers struct {
	WorkflowHandlers      *workflow_handlers.WorkflowHandlers
	WalletHandlers        *wallet_handlers.WalletHandlers
	VirtualCardHandlers   *virtualcard_handlers.VirtualCardHandlers
	UtilizationHandlers   *utilization_handlers.UtilizationHandlers
	UserHandlers          *user_handlers.UserHandlers
	TransactionHandlers   *transaction_handlers.TransactionHandlers
	TillHandlers          *till_handlers.TillHandlers
	TaxHandlers           *tax_handlers.TaxHandlers
	SupplierHandlers      *supplier_handlers.SupplierHandlers
	ShippingHandlers      *shipping_handlers.ShippingHandlers
	ShipperHandlers       *shipper_handlers.ShipperHandlers
	SellHandlers          *sell_handlers.SellHandlers
	SaleHandlers          *sale_handlers.SaleHandlers
	ReportHandlers        *report_handlers.ReportHandlers
	RegisterHandlers      *register_handlers.RegisterHandlers
	ReferenceHandlers     *reference_handlers.ReferenceHandlers
	PurchaseHandlers      *purchase_handlers.PurchaseHandlers
	PromotionHandlers     *promotion_handlers.PromotionHandlers
	ProductHandlers       *product_handlers.ProductHandlers
	PaymentHandlers       *payment_handlers.PaymentHandlers
	OrderHandlers         *order_handlers.OrderHandlers
	MemberHandlers        *member_handlers.MemberHandlers
	MarketingHandlers     *marketing_handlers.MarketingHandlers
	InvoiceHandlers       *invoice_handlers.InvoiceHandlers
	InventoryHandlers     *inventory_handlers.InventoryHandlers
	EodHandlers           *eod_handlers.EodHandlers
	EmployeeHandlers      *employee_handlers.EmployeeHandlers
	DepartmentHandlers    *department_handlers.DepartmentHandlers
	DependantHandlers     *dependant_handlers.DependantHandlers
	DatalogHandlers       *datalog_handlers.DatalogHandlers
	CustomerHandlers      *customer_handlers.CustomerHandlers
	ContactHandlers       *contact_handlers.ContactHandlers
	ConsumptionHandlers   *consumption_handlers.ConsumptionHandlers
	ChangeHandlers        *change_handlers.ChangeHandlers
	CatalogueHandlers     *catalogue_handlers.CatalogueHandlers
	AuthHandlers          *auth_handlers.AuthHandlers
	AppStoreHandlers      *appstore_handlers.AppStoreHandlers
	AppHandlers           *app_handlers.AppHandlers
	StoreHandlers         *store_handlers.StoreHandlers
	ChannelHandlers       *channel_handlers.ChannelHandlers
	PaymentMethodHandlers *payment_method_handlers.PaymentMethodHandlers
	GroupHandlers         *group_handlers.GroupHandlers
	AclHandlers           *acl_handlers.AclHandlers
	RoleHandlers          *role_handlers.RoleHandlers
	PermissionHandlers    *permission_handlers.PermissionHandlers
}

func New() *Handlers {

	return &Handlers{
		WorkflowHandlers:      workflow_handlers.New(),
		WalletHandlers:        wallet_handlers.New(),
		VirtualCardHandlers:   virtualcard_handlers.New(),
		UtilizationHandlers:   utilization_handlers.New(),
		UserHandlers:          user_handlers.New(),
		TransactionHandlers:   transaction_handlers.New(),
		TillHandlers:          till_handlers.New(),
		TaxHandlers:           tax_handlers.New(),
		SupplierHandlers:      supplier_handlers.New(),
		ShippingHandlers:      shipping_handlers.New(),
		ShipperHandlers:       shipper_handlers.New(),
		SellHandlers:          sell_handlers.New(),
		SaleHandlers:          sale_handlers.New(),
		ReportHandlers:        report_handlers.New(),
		RegisterHandlers:      register_handlers.New(),
		ReferenceHandlers:     reference_handlers.New(),
		PurchaseHandlers:      purchase_handlers.New(),
		PromotionHandlers:     promotion_handlers.New(),
		ProductHandlers:       product_handlers.New(),
		PaymentHandlers:       payment_handlers.New(),
		OrderHandlers:         order_handlers.New(),
		MemberHandlers:        member_handlers.New(),
		MarketingHandlers:     marketing_handlers.New(),
		InvoiceHandlers:       invoice_handlers.New(),
		InventoryHandlers:     inventory_handlers.New(),
		EodHandlers:           eod_handlers.New(),
		EmployeeHandlers:      employee_handlers.New(),
		DepartmentHandlers:    department_handlers.New(),
		DependantHandlers:     dependant_handlers.New(),
		DatalogHandlers:       datalog_handlers.New(),
		CustomerHandlers:      customer_handlers.New(),
		ContactHandlers:       contact_handlers.New(),
		ConsumptionHandlers:   consumption_handlers.New(),
		ChangeHandlers:        change_handlers.New(),
		CatalogueHandlers:     catalogue_handlers.New(),
		AuthHandlers:          auth_handlers.New(),
		AppStoreHandlers:      appstore_handlers.New(),
		AppHandlers:           app_handlers.New(),
		StoreHandlers:         store_handlers.New(),
		ChannelHandlers:       channel_handlers.New(),
		PaymentMethodHandlers: payment_method_handlers.New(),
		GroupHandlers:         group_handlers.New(),
		AclHandlers:           acl_handlers.New(),
		RoleHandlers:          role_handlers.New(),
		PermissionHandlers:    permission_handlers.New(),
	}

}
