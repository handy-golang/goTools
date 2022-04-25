package main

import (
	"fmt"
	"time"

	"github.com/EasyGolang/goTools/mEncrypt"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	SecretKey := "meichangliang"

	tokenStr1 := mEncrypt.NewToken(mEncrypt.NewTokenOpt{
		SecretKey: SecretKey,                        // key
		ExpiresAt: time.Now().Add(60 * time.Second), // 过期时间
		Message:   "kjhgfyu",
		Issuer:    "mo7.cc",
		Subject:   "UserToken",
	}).Generate()

	tokenStr2 := mEncrypt.NewToken(mEncrypt.NewTokenOpt{
		SecretKey: SecretKey,                        // key
		ExpiresAt: time.Now().Add(60 * time.Second), // 过期时间
		Message:   "abcde",
		Issuer:    "mo7.cc",
		Subject:   "UserToken",
	}).Generate()

	fmt.Println("tokenStr1", tokenStr1)
	fmt.Println("tokenStr2", tokenStr2)
	mes2 := mEncrypt.ParseToken(tokenStr2)

	mes1 := mEncrypt.ParseToken(tokenStr1)

	fmt.Println("mes1", mes1)
	fmt.Println("mes2", mes2)

	fmt.Println(" =========   END   ========= ")
}
