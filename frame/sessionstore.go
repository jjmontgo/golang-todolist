package frame

import (
	"os"
	"github.com/srinathgs/mysqlstore"
)

var SessionStore *mysqlstore.MySQLStore
var sessionErr error
var sessionStoreInitialized bool

func GetSessionStore() *mysqlstore.MySQLStore {
	if sessionStoreInitialized == true {
		return SessionStore
	}

	SessionStore, sessionErr = mysqlstore.NewMySQLStore(
		os.Getenv("SESSION_MYSQL_USERNAME") + ":" +
		os.Getenv("SESSION_MYSQL_PASSWORD") + "@tcp(" +
		os.Getenv("SESSION_MYSQL_HOST") + ":3306)/" +
		os.Getenv("SESSION_MYSQL_DB") + "?parseTime=true&loc=Local",
		os.Getenv("SESSION_MYSQL_TABLE"),
		"/",
		3600,
		[]byte("gewdewdsedfs"))
	if sessionErr != nil {
		panic(sessionErr)
	}
	sessionStoreInitialized = true
	return SessionStore
}
