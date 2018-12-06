package frame

import (
	"log"
	"os"
	"github.com/gorilla/sessions"
	"github.com/savaki/dynastore"
	"github.com/srinathgs/mysqlstore"
)

var SessionStore sessions.Store
var sessionErr error
var sessionStoreInitialized bool

func GetSessionStore() sessions.Store {
	if sessionStoreInitialized == true {
		return SessionStore
	}

	mode := os.Getenv("MODE")
	if mode == "prod" {
		useDynamoDbSessionStore()
	} else {
		useMysqlSessionStore()
	}

	sessionStoreInitialized = true
	return SessionStore
}

func useMysqlSessionStore() {
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
	  log.Fatalln(sessionErr)
	}
}

func useDynamoDbSessionStore() {
	SessionStore, sessionErr = dynastore.New(dynastore.Path("/"), dynastore.HTTPOnly())
	if sessionErr != nil {
	  log.Fatalln(sessionErr)
	}
}
