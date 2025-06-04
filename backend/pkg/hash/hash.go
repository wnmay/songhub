package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hashed, password string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}