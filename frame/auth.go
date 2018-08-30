package frame

import (
	"log"
	"golang.org/x/crypto/bcrypt"
)

func UserIsLoggedIn() bool {
	return SessionGet("username") != ""
}

func HashPassword(password string) string {
	passwordByteSlice := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(passwordByteSlice, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func VerifyPassword(password string, hash string) bool {
	passwordByteSlice := []byte(password)
	hashByteSlice := []byte(hash)
	err := bcrypt.CompareHashAndPassword(hashByteSlice, passwordByteSlice)
	if err != nil {
		return false
	}
	return true
}
