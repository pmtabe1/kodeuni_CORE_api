package pos_databases_migrations

import (
	"log"

	"github.com/paulmsegeya/pos/core/models/auth_models"
	"github.com/paulmsegeya/pos/core/models/pos_models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	log.Println("about to initiated  migration ....")
	log.Println(db)

	if db == nil {
		log.Panicln("DB INSTANCE IS NIL it should't be")
	}

	log.Println(" Executing Core Models DB Migrations .....")

	err := db.AutoMigrate(&pos_models.Employee{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Consumption{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Channel{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Report{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.App{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Promotion{})
	log.Println(err)
	err = db.AutoMigrate(&pos_models.Product{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.ProductItem{})
	log.Println(err)
	err = db.AutoMigrate(&pos_models.Marketing{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Utilization{})
	log.Println(err)

	err = db.AutoMigrate(&auth_models.Profile{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Eod{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Tax{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.TaxItem{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Department{})
	log.Println(err)
	err = db.AutoMigrate(&pos_models.Invoice{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.InvoiceItem{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Purchase{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.PurchaseItem{})
	log.Println(err)
	// err = db.AutoMigrate(&pos_models.Workflow{})
	// log.Println(err)

	err = db.AutoMigrate(&pos_models.Spending{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Wallet{})
	log.Println(err)
	err = db.AutoMigrate(&pos_models.Register{})
	log.Println(err)
	err = db.AutoMigrate(&pos_models.Change{})
	log.Println(err)
	// err = db.AutoMigrate(&pos_models.InvoiceItem{})
	// log.Println(err)
	err = db.AutoMigrate(&pos_models.Transaction{})
	log.Println(err)
	err = db.AutoMigrate(&pos_models.Member{})
	log.Println(err)
	err = db.AutoMigrate(&pos_models.Dependant{})
	log.Println(err)
	err = db.AutoMigrate(&pos_models.Purchase{})
	log.Println(err)
	err = db.AutoMigrate(&pos_models.PurchaseItem{})
	log.Println(err)
	err = db.AutoMigrate(&auth_models.User{})
	log.Println(err)
	err = db.AutoMigrate(&auth_models.Acl{})
	log.Println(err)
	err = db.AutoMigrate(&auth_models.Role{})
	log.Println(err)
	err = db.AutoMigrate(&auth_models.Permission{})
	log.Println(err)
	err = db.AutoMigrate(&auth_models.Group{})
	log.Println(err)
	err = db.AutoMigrate(&pos_models.Order{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Datalog{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Customer{})
	log.Println(err)
	err = db.AutoMigrate(&pos_models.Catalogue{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Change{})
	log.Println(err)
	err = db.AutoMigrate(&pos_models.VirtualCard{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.AppStore{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.OrderItem{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Reference{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Tax{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.TaxItem{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Store{})
	log.Println(err)
	err = db.AutoMigrate(&pos_models.Secret{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Supplier{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.PaymentMethod{})
	log.Println(err)
	err = db.AutoMigrate(&pos_models.Sale{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.SaleItem{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Order{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Contact{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Store{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Inventory{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Receipt{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Till{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Register{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Shipping{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Shipper{})
	log.Println(err)

	err = db.AutoMigrate(&pos_models.Workflow{})
	log.Println(err)
	if err != nil {
		log.Panicln("Migration Error" + err.Error())
	}

	return err
}
