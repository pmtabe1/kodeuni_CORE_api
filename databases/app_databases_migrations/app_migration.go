package pos_databases_migrations

import (
	"log"

	"github.com/paulmsegeya/subscription/core/models/auth_models"
	"github.com/paulmsegeya/subscription/core/models/data_models"
	"github.com/paulmsegeya/subscription/core/models/subscription_models"
	"github.com/paulmsegeya/subscription/core/models/workflow_models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	log.Println("about to initiated  migration ....")
	log.Println(db)

	if db == nil {
		log.Panicln("DB INSTANCE IS NIL it should't be")
	}

	log.Println(" Executing Core Models DB Migrations .....")

	err := db.AutoMigrate(&subscription_models.Staff{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Team{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Agreement{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Report{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Customer{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Department{})
	log.Println(err)
	err = db.AutoMigrate(&subscription_models.Product{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Subscriber{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Subscription{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Service{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Schedule{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Licence{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Limitation{})
	log.Println(err)

	err = db.AutoMigrate(&auth_models.Profile{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Department{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Notification{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Licence{})
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

	err = db.AutoMigrate(&auth_models.Secret{})
	log.Println(err)

	err = db.AutoMigrate(&data_models.Datalog{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Customer{})
	log.Println(err)
	err = db.AutoMigrate(&subscription_models.Schedule{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Service{})
	log.Println(err)
	err = db.AutoMigrate(&subscription_models.Staff{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Product{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Contract{})
	log.Println(err)

	err = db.AutoMigrate(&subscription_models.Contact{})
	log.Println(err)

	err = db.AutoMigrate(&workflow_models.Workflow{})
	log.Println(err)
	if err != nil {
		log.Panicln("Migration Error" + err.Error())
	}

	return err
}
