package app_databases

import (
	"fmt"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Sannity() {

	// github.com/denisenkom/go-mssqldb
	//sqlserver://username:password@host/instance?param1=value&param2=value
	//dsn := "sqlserver://sa:yatch123@SAGESERVER:1433?database=YATCH_INTEGRATION"
	//db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	dsn := "root:gr00t00@tcp(127.0.0.1:3306)/YATCH_INTEGRATION?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dbState := false
	dbConn, err := db.DB()

	if err != nil {
		dbState = false
		fmt.Println(err)
		//os.Exit(1)
		fmt.Println(err.Error())
	}

	err = dbConn.Ping()

	if err != nil {
		dbState = false
		fmt.Println("PING err:" + err.Error())
		//os.Exit(1)
		fmt.Println(err.Error())
	} else {

		fmt.Println("SUCCESSFULLY PINGED  DATABASE  --- IT IS ALIVE")
		dbState = true
		fmt.Println(dbState)
	}

	fmt.Println(err)
	fmt.Println(db)

}
