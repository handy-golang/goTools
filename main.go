package main

import (
	"fmt"

	"github.com/EasyGolang/goTools/mMongo"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	db := mMongo.New(mMongo.Opt{
		UserName: "mo7",
		Password: "asdasd55555",
		Host:     "mo7.cc",
		Port:     "17017",
	})
	db.Connect()
	db.Ping()

	db.Close()

	fmt.Println(" =========   END   ========= ")
}
