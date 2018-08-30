package frame

import (
	"log"
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
	//defer store.Close()

	return SessionStore
}

func SessionSet(field string, value interface{}) {
	session, err := GetSessionStore().Get(Registry.Request, os.Getenv("SESSION_NAME"))
	session.Values[field] = value
	err = session.Save(Registry.Request, Registry.Response)
	if err != nil {
		log.Fatalf("SessionSet(): %q\n", err)
	}
}

func SessionGet(field string) interface{} {
	session, err := GetSessionStore().Get(Registry.Request, os.Getenv("SESSION_NAME"))
	if err != nil {
		log.Fatalf("SessionGet(): %q\n", err)
	}
	return session.Values[field]
}
