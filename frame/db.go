package frame

import (
	"upper.io/db.v2/mysql"
	"upper.io/db.v2/lib/sqlbuilder"
	"log"
	"os"
)

var Db sqlbuilder.Database
var dbErr error
var dbInitialized bool

func DB() sqlbuilder.Database {
	if dbInitialized == true {
		return Db
	}

	var settings = mysql.ConnectionURL{
		Host: os.Getenv("MYSQL_HOST"),
		User: os.Getenv("MYSQL_USERNAME"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Database: os.Getenv("MYSQL_DB")}

	Db, dbErr = mysql.Open(settings)
	if dbErr != nil {
		log.Fatalf("DB(): %q", dbErr)
		return nil
	}

	return Db
}
