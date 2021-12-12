package pos_databases

import (
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/paulmsegeya/pos/config/pos_config"
	mysql "gorm.io/driver/mysql"

	sqlserver "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	MYSQL = "mysql"
	MSSQL = "mssql"
	FROM  = "from"
	TO    = "to"
)

type IPosDatabase interface {
}

type PosDatabase struct {
	*pos_config.PosConfig
}

func New() *PosDatabase {

	jsonConfigString := pos_config.New().LoadJSONConfig("")
	configuration := pos_config.New().FromJSON(jsonConfigString)
	posConfig := configuration.LoadConfiguration()

	return &PosDatabase{
		PosConfig: posConfig,
	}
}

func (s *PosDatabase) PingDatabaseStatus(dbInstance *gorm.DB) bool {
	dbState := false

	if dbInstance == nil {
		log.Println("GORM DB INSTANCE IS NIL ")
		return false
	} else {

		dbConn, err := dbInstance.DB()

		if err != nil {
			dbState = false
			fmt.Println(err)
			log.Println(err.Error())

			return dbState
		}

		err = dbConn.Ping()

		if err != nil {
			dbState = false
			fmt.Println("PING err:" + err.Error())
			log.Println(err.Error())
		} else {

			log.Println("SUCCESSFULLY PINGED  DATABASE  --- IT IS ALIVE")
			dbState = true
		}
		log.Println(s.PosConfig)

		switch s.Integration.RouteSource {
		case TO:
			if dbState {

				log.Println("Successfully Connected to database  ==>" + s.PosConfig.Integration.To.Database.DBName + " ==> dbengine:" + s.PosConfig.Integration.To.Database.Engine)
			} else {
				log.Println("Failed to Connect to database  ==>" + s.PosConfig.Integration.To.Database.DBName + " ==> dbengine:" + s.PosConfig.Integration.To.Database.Engine)

			}

		case FROM:
			if dbState {

				log.Println("Successfully Connected to database  ==>" + s.PosConfig.Integration.From.Database.DBName + " ==> dbengine:" + s.PosConfig.Integration.From.Database.Engine)
			} else {
				log.Println("Failed to Connect to database  ==>" + s.PosConfig.Integration.From.Database.DBName + " ==> dbengine:" + s.PosConfig.Integration.From.Database.Engine)

			}

		default:
			log.Println("Unknown Ping database config detected...")
		}

	}

	return dbState
}

func (s *PosDatabase) LoadConfigulationSettings() *PosDatabase {

	jsonConfigString := s.LoadJSONConfig("")
	configuration := pos_config.New().FromJSON(jsonConfigString)
	s.PosConfig = configuration.LoadConfiguration()

	if s.PosConfig == nil {
		s.LoadConfiguration()
	}

	return s
}

func (s *PosDatabase) DBConnection() (dInstance *gorm.DB) {
	//Load settings
	s.LoadConfigulationSettings()

	if s.PosConfig == nil {
		log.Panicln("RouteSource is EMPTy")
	}

	if len(s.Integration.RouteSource) == 0 {
		log.Panicln("Route is not provided please specify")
	}

	switch s.Integration.RouteSource {
	case FROM:
		log.Println(s.Integration.RouteSource)
		dInstance = s.ConnectFromDB()
	case TO:
		dInstance = s.ConnectToDB()
	}

	log.Println(dInstance)
	return dInstance
}

func (s *PosDatabase) ConnectFromDB() *gorm.DB {

	// Load all config stuff here

	if s.PosConfig == nil {
		s.PosConfig = s.LoadConfiguration()
		log.Println("ERROR sage configuration must never be nil")
	}

	var dbInstance *gorm.DB
	log.Println(s.PosConfig.Integration.From.Database.Engine)

	gormConfig := &gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   nil,
		DryRun:                                   false,
		PrepareStmt:                              true,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		ClauseBuilders:                           map[string]clause.ClauseBuilder{},
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  map[string]gorm.Plugin{},
	}

	switch s.PosConfig.Integration.From.Database.Engine {
	case MYSQL:

		connectionString := s.ConnectionStringFrom(pos_config.New().ReadConfiguration())

		db, err := gorm.Open(mysql.Open(connectionString), gormConfig)

		//	log.Println("DB Connection ERROR " + err.Error())
		if err != nil {
			message := fmt.Sprintf("Error connecting to database using GORM : error=%v", err)
			//log.Println(message)
			log.Panicln(message)
			//return nil, nil, nil
		}

		// db, err := gorm.Open("mssql", ConnectionString)
		// if err != nil {
		// 	log.Panicln("Failed Connection to Source Database")
		// }

		dbInstance = db

		//err:=db.DB()
		//err = db.DB().Ping()

		if err != nil {
			log.Panicln(" Error :" + err.Error())
		} else {
			s.PingDatabaseStatus(dbInstance)
		}
		return db

	case MSSQL:

		connectionString := s.ConnectionStringFrom(pos_config.New().ReadConfiguration())

		db, err := gorm.Open(sqlserver.Open(connectionString), gormConfig)

		if err != nil {
			log.Println("DB Connection ERROR " + err.Error())
			message := fmt.Sprintf("Error connecting to database using GORM : error=%v", err)
			//log.Println(message)
			log.Panicln(message)
			//return nil, nil, nil
		}

		if err != nil {

			log.Panicln("Failed Connection to Source Database")
		}

		// err = db.DB().Ping()

		if err != nil {
			log.Panicln(" Error :" + err.Error())
		} else {
			s.PingDatabaseStatus(dbInstance)
		}

		dbInstance = db

		return db

	}

	if dbInstance == nil {
		log.Println("DB INSTANCE IS NIL RECREATING IT ")
		dbInstance = s.DBConnection()
	}

	return dbInstance

}

func (s *PosDatabase) ConnectToDB() *gorm.DB {
	log.Println(s.PosConfig.Integration.RouteSource)
	log.Println(s.PosConfig.Integration.To.Database.DBName)

	if s.PosConfig == nil {
		s.PosConfig = s.LoadConfiguration()
		log.Println("ERROR sage configuration must never be nil")
	}

	var dbInstance *gorm.DB
	log.Println(">>>>>>>")
	log.Println(s.PosConfig.Integration.To.Database.Engine)

	gormConfig := &gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   nil,
		DryRun:                                   false,
		PrepareStmt:                              true,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		ClauseBuilders:                           map[string]clause.ClauseBuilder{},
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  map[string]gorm.Plugin{},
	}

	switch s.PosConfig.Integration.To.Database.Engine {
	case MYSQL:

		connectionString := s.ConnectionStringTo(pos_config.New().ReadConfiguration())

		log.Println(connectionString)

		if len(connectionString) == 0 {
			log.Panicln("MYSQL CONNECTION STRING MISSING")
		}

		db, err := gorm.Open(mysql.Open(connectionString), gormConfig)

		if err != nil {
			log.Println("DB Connection ERROR " + err.Error())
			message := fmt.Sprintf("Error connecting to database using GORM : error=%v", err)
			log.Println(message)
			//return nil, nil, nil
		}

		dbInstance = db

		if err != nil {
			log.Println("Ping Error :" + err.Error())
		} else {

			s.PingDatabaseStatus(dbInstance)
		}
		return db

	case MSSQL:

		connectionString := s.ConnectionStringTo(pos_config.New().ReadConfiguration())
		log.Println(connectionString)

		if len(connectionString) == 0 {
			log.Panicln("MYSQL CONNECTION STRING MISSING")
		}

		db, err := gorm.Open(sqlserver.Open(connectionString), gormConfig)

		//db, err := gorm.Open("mssql", connectionString)

		if err != nil {
			log.Println("DB Connection ERROR " + err.Error())
			message := fmt.Sprintf("Error connecting to database using GORM : error=%v", err)
			log.Println(message)
			//log.Panicln(message)
			//return nil, nil, nil
		}

		//db, err := gorm.Open("mysql", ConnectionString)

		if err != nil {
			log.Println("Failed Connection to Source Database due to :" + err.Error())
		}

		// err = db.DB().Ping()

		dbInstance = db

		if err != nil {
			log.Println("Ping Error :" + err.Error())
		} else {
			s.PingDatabaseStatus(dbInstance)

		}

		return db

	}

	return dbInstance
}

func (s *PosDatabase) ConnectionStringFrom(configuration pos_config.PosConfig) (connString string) {

	log.Println(configuration)
	s.PosConfig = &configuration

	var port string

	switch configuration.Integration.From.Database.Engine {
	case MSSQL:
		if len(configuration.Integration.From.Database.Port) == 0 {
			port = "1433" //DEFAULT PORT mssql
		} else {
			port = configuration.Integration.From.Database.Port
		}
		//dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
		//  db, err := gorm.Open("mssql", "sqlserver://username:password@localhost:1433?

		connString = fmt.Sprintf("sqlserver://%v:%v@%v:%v?database=%v",
			configuration.Integration.From.Database.Username, configuration.Integration.From.Database.Password, configuration.Integration.From.Database.Host, port, configuration.Integration.From.Database.DBName)

	case MYSQL:

		if len(configuration.Integration.From.Database.Port) == 0 {
			port = "3306" //DEFAULT PORT mysql
		} else {
			port = configuration.Integration.From.Database.Port
		}
		connString = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			configuration.Integration.From.Database.Username, configuration.Integration.From.Database.Password, configuration.Integration.From.Database.Host, port, configuration.Integration.From.Database.DBName)

	}

	return connString

}

func (s *PosDatabase) ConnectionStringTo(configuration pos_config.PosConfig) (connString string) {
	log.Println(configuration)
	s.PosConfig = &configuration

	var port string

	switch configuration.Integration.To.Database.Engine {
	case MSSQL:
		if len(configuration.Integration.To.Database.Port) == 0 {
			port = "1433" //DEFAULT PORT mssql
		} else {
			port = configuration.Integration.To.Database.Port
		}
		//dsn := "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"

		connString = fmt.Sprintf("sqlserver://%v:%v@%v:%v?database=%v",
			configuration.Integration.To.Database.Username, configuration.Integration.To.Database.Password, configuration.Integration.To.Database.Host, port, configuration.Integration.To.Database.DBName)

	case MYSQL:

		if len(configuration.Integration.To.Database.Port) == 0 {
			port = "3306" //DEFAULT PORT mysql
		} else {
			port = configuration.Integration.To.Database.Port
		}
		connString = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			configuration.Integration.To.Database.Username, configuration.Integration.To.Database.Password, configuration.Integration.To.Database.Host, port, configuration.Integration.To.Database.DBName)

	}

	return connString

}
