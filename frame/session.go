package frame

import (
	"log"
	"os"
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
	//defer store.Close()

	return SessionStore
}

func SessionSetVar(field string, value interface{}) {
	session, err := GetSessionStore().Get(Registry.Request, os.Getenv("SESSION_NAME"))
	session.Values[field] = value
	err = session.Save(Registry.Request, Registry.Response)
	if err != nil {
		log.Fatalf("SessionSetVar(): %q\n", err)
	}
}

func SessionGetVar(field string) interface{} {
	session, err := GetSessionStore().Get(Registry.Request, os.Getenv("SESSION_NAME"))
	if err != nil {
		log.Fatalf("SessionGetVar(): %q\n", err)
	}
	return session.Values[field]
}

// func SessionGetUser() *model.User {
// 	val := SessionGetVar("user")
// 	var user = &model.User{}
// 	if user, ok := val.(*model.User); !ok {
// 		return nil
// 	}
// 	return user
// }

func UserIsLoggedIn() bool {
	return SessionGetVar("user") != nil
}
