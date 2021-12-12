package pos_databases_migrations

import (
	"log"
	//"os"
	"testing"

	"github.com/paulmsegeya/pos/databases/pos_databases"
	"github.com/stretchr/testify/require"
)

func TestMigrateMysql(t *testing.T) {
	//os.Setenv("SAGE_CONFIG", "./config.%v.json")
	database := pos_databases.New()
	gormInstance := database.ConnectToDB()

	if gormInstance == nil {
		log.Println("DB Instance is Null")
	}
	err := Migrate(gormInstance)
	require.Nilf(t, err, "Expected error to  nil but got %v instead", err)

}

func TestMigrateMssql(t *testing.T) {
	//os.Setenv("SAGE_CONFIG", "./config.%v.json")

	database := pos_databases.New()
	gormInstance := database.ConnectFromDB()

	if gormInstance == nil {
		log.Println("DB Instance is Null")
	}
	err := Migrate(gormInstance)
	require.Nilf(t, err, "Expected error to  nil but got %v instead", err)

}
