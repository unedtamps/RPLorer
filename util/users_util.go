package util

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenereateHasedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}

func CompareHashedPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func GenerateUUID() string {
	return uuid.NewString()
}
