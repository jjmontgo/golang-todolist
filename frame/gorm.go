package frame

import (
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var gormDB *gorm.DB
var dbErr error
var dbInitialized bool

func GORM() *gorm.DB {
	if dbInitialized == true {
		return gormDB
	}

	host := os.Getenv("MYSQL_HOST")
	user := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	database := os.Getenv("MYSQL_DB")
	dsn :=  user + ":" + password + "@tcp(" + host + ":3306)/" + database + "?charset=utf8&parseTime=True&loc=Local"
	gormDB, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("Failed to connect database")
	}

	gormDB.SingularTable(true) // singular table names

	if os.Getenv("MODE") == "dev" {
		gormDB.LogMode(true)
	}

	return gormDB
}
