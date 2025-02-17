package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashingPassword(paswd string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(paswd), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error while hashing the password: ", err)
		return "", err
	}
	return string(bytes), nil
}

func ValidatePassword(hashedPaswd string, plainPaswd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPaswd), []byte(plainPaswd))
	if err != nil {
		fmt.Println("Error while comparing the password: ", err)
		return false
	}
	return true
}


