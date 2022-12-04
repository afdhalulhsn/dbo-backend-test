package connection

import (
	"dbo_backend_test/app/domain/entity"
	"dbo_backend_test/internal/config"
	"dbo_backend_test/internal/config/db"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var (
	DbConnection *gorm.DB
)

var table = []interface{}{
	&entity.LoginHistory{},
	&entity.DataOrder{},
	&entity.UserData{},
	&entity.CustomerData{},
	&entity.ProductData{},
}

func init() {
	var err error
	fmt.Printf("Init DB Connection =>")
	cfg := config.GetConfig()
	DbConnection, err = Conn(cfg.Database.Dbo)
	if err != nil {
		panic("Cannot Connect to Database")
	}
	MigrateSchema(DbConnection, table)
}

func Conn(cfg db.Database) (*gorm.DB, error) {
	fmt.Printf("Connect to Database %v", cfg.Dbname)
	user := cfg.Username
	password := cfg.Password
	host := cfg.Host
	port := cfg.Port
	dbname := cfg.Dbname

	// fmt.Println(user + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local")
	dsn := user + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	// dbConn, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	conn, err := dbConn.DB()
	if err != nil {
		return nil, err
	}
	conn.SetMaxIdleConns(7)
	conn.SetMaxOpenConns(10)
	conn.SetConnMaxLifetime(1 * time.Hour)

	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func MigrateSchema(db *gorm.DB, tables []interface{}) {
	err := db.AutoMigrate(tables...)
	if err != nil {
		log.Fatalln("Error While Migrateing Schema", err.Error())
	}
}
