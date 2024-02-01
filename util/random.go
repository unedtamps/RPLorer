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

func RandomTypePremium() string {
	types := []string{
		"dcaa02fb-e960-40db-b38b-75dc104fb017",
		"59318f8a-3e06-49e7-86f7-006039fb2112",
	}
	return types[rand.Intn(len(types))]
}
