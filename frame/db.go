package frame

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var connectionEstablished bool
var db *sql.DB
var err error

func DB() *sql.DB {
	if connectionEstablished == true {
		return db
	}
	user := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DB")
	host := os.Getenv("MYSQL_HOST")
	db, err = sql.Open("mysql", user+":"+password+"@tcp("+host+")/"+dbname)
	if err != nil {
		panic(err.Error())
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	connectionEstablished = true;

	return db
}

