package mMongo

import (
	"context"
	"fmt"
	"strings"

	"github.com/EasyGolang/goTools/mStr"
	"go.mongodb.org/mongo-driver/mongo"
)

type Opt struct {
	URI      string
	UserName string
	Password string
	Address  string
	DBName   string
	Timeout  int // 秒
	Event    func(string, string)
}

type DB struct {
	URI     string
	Event   func(string, string)
	Client  *mongo.Client
	dbName  string
	Ctx     context.Context
	cancel  context.CancelFunc
	db      *mongo.Database
	Table   *mongo.Collection
	Timeout int // 超时时长
}

func New(opt Opt) *DB {
	var optNilStr []string
	if len(opt.DBName) < 2 {
		optNilStr = append(optNilStr, "Database")
	}
	// 如果没有 URI 则要检查 账户信息
	if len(opt.URI) < 5 {
		switch {
		case len(opt.UserName) < 2:
			optNilStr = append(optNilStr, "UserName")
			fallthrough
		case len(opt.Password) < 2:
			optNilStr = append(optNilStr, "Password")
			fallthrough
		case len(opt.Address) < 3:
			optNilStr = append(optNilStr, "Address")
		}
	}

	if len(optNilStr) > 0 {
		fmt.Println("缺少参数:", strings.Join(optNilStr, ","))
	}
	var NewDB DB

	NewDB.dbName = opt.DBName

	if opt.Timeout < 1 {
		NewDB.Timeout = 60
	} else {
		NewDB.Timeout = opt.Timeout
	}

	if len(opt.URI) > 5 {
		NewDB.URI = opt.URI
	} else {
		/*

			mongosh "mongodb://root:asdasd55555@mo7.cc:17017/Hunter?authSource=Hunter"

		*/
		NewDB.URI = mStr.Join(
			"mongodb://",
			opt.UserName, ":", opt.Password,
			"@", opt.Address,
			"/", NewDB.dbName,
		)
	}

	if opt.Event != nil {
		NewDB.Event = opt.Event
	} else {
		NewDB.Event = func(s1, s2 string) {}
	}

	return &NewDB
}
