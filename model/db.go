package model

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// get config from .env when run locally
	// localEnvFile "github.com/joho/godotenv"
)

var connectionEstablished bool
var db *sql.DB
var err error

func DB() *sql.DB {
	if connectionEstablished == true {
		return db
	}
	db, err = sql.Open("mysql", "root:qazqaz@/golang_todolist")
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

