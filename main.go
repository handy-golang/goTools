package main

import (
	_ "embed"
	"fmt"
	"time"

	"github.com/EasyGolang/goTools/mEncrypt"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	token := mEncrypt.NewToken(mEncrypt.NewTokenOpt{
		SecretKey: "abc",
		ExpiresAt: time.Now().Add(time.Hour),
		Message:   "墨七太帅了",
		Issuer:    "mo7.cc",
		Subject:   ".net",
	}).Generate()

	fmt.Println(token)

	msg, err := mEncrypt.ParseToken(token, "abc")
	fmt.Println(msg.Message)
	fmt.Println(err)

	fmt.Println(" =========   END   ========= ")
}
