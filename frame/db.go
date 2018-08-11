package frame

import (
	"upper.io/db.v2/mysql"
	"upper.io/db.v2/lib/sqlbuilder"
	"os"
	"log"
)

var Db sqlbuilder.Database
var err error
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

	Db, err = mysql.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}

	return Db
}
