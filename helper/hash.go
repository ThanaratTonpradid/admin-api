package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(plainPwd string) (string, error) {
	bytePlainPwd := []byte(plainPwd)
	hash, err := bcrypt.GenerateFromPassword(bytePlainPwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CompareHashPassword(hashedPwd, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePlainPwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlainPwd)
	return err == nil
}
