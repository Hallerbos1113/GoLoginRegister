package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	pwd := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(pwd, 14)

	if err != nil {
		log.Fatal(err)
	}

	return string(hash)
}

func CheckPassword(hpass string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hpass), []byte(pass))
	return err == nil
}

func GenerateUserID(uid string) string {
	return uid
}