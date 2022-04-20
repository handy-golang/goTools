package main

import (
	"fmt"

	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mStr"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	db := mMongo.New(mMongo.Opt{
		UserName: "mo7",
		Password: "asdasd55555",
		Host:     "mo7.cc",
		Port:     "17017",
		DBName:   "Hunter",
	}).Connect().Collection("HotList")

	cursor, err := db.Table.Find(db.Ctx, bson.D{{}})

	fmt.Println(mStr.ToStr(cursor), err)

	db.Close()

	fmt.Println(" =========   END   ========= ")
}
