package util

import (
	"math/rand"
	"time"
)

const alphanumeric string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func CheckErr(err error) bool {
	if err != nil {
		return true
	}
	return false
}

func RandomBoolean() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func RandString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = alphanumeric[rand.Intn(len(alphanumeric))]
	}
	return string(b)
}

func Ptr[T any](t T) *T {
	return &t
}
