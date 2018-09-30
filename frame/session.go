package frame

import (
	"log"
	"os"
	"github.com/gorilla/sessions"
	"github.com/srinathgs/mysqlstore"
	// "golang-todolist/model"
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
	return SessionStore
}

func SessionSetVar(field string, value interface{}) {
	session := GetSession()
	session.Values[field] = value
}

func SessionGetVar(field string) interface{} {
	session := GetSession()
	return session.Values[field]
}

func GetSession() *sessions.Session {
	session, err := GetSessionStore().Get(Registry.Request, os.Getenv("SESSION_NAME"))
	if err != nil {
		log.Fatalf("GetSession(): %q\n", err)
	}
	return session
}

// called by dispatch.go
func SessionSave() {
	session := GetSession()
	err := session.Save(Registry.Request, Registry.Response)
	if err != nil {
		log.Fatalf("SessionSave(): %q\n", err)
	}
}

func UserIsLoggedIn() bool {
	return SessionGetVar("user") != nil
}
