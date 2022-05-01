package mEncrypt

import (
	"math/rand"
	"time"
)

func RandStr(length int) string {
	baseStr := "0123456789abcdefghijklmnopqrstuvwxyz"

	r := rand.New(rand.NewSource(time.Now().UnixNano() + rand.Int63()))
	bytes := make([]byte, length)
	l := len(baseStr)
	for i := 0; i < length; i++ {
		bytes[i] = baseStr[r.Intn(l)]
	}
	return string(bytes)
}
