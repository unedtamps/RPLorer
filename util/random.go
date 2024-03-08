package util

import (
	"math/rand"

	"github.com/google/uuid"
)

func RandomUUID() string {
	uuid := uuid.NewString()
	return uuid
}

var alphabets = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString() string {
	random := make([]rune, 15)
	len_alphabets := len(alphabets)
	for i := range random {
		random[i] = alphabets[rand.Intn(len_alphabets)]
	}
	return string(random)
}

func RandomEmail() string {
	return RandomString() + "@gmail.com"
}

func RandomHashedPassword() string {
	randomPass := RandomString()
	hash, _ := GenereateHasedPassword(randomPass)
	return string(hash)
}

func RandomInt64() int64 {
	return rand.Int63()
}
func RandomInt8() int {
	return rand.Intn(10)
}
