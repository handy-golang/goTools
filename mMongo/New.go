package mMongo

import (
	"fmt"
	"strings"
)

type Opt struct {
	UserName   string
	Password   string
	Host       string
	Port       string
	Database   string
	Collection string
	Event      func(string, string)
}

type DB struct{}

func New(opt Opt) *DB {
	var optNilStr []string
	switch {
	case len(opt.UserName) < 2:
		optNilStr = append(optNilStr, "UserName")
		fallthrough
	case len(opt.Password) < 2:
		optNilStr = append(optNilStr, "Password")
		fallthrough
	case len(opt.Host) < 2:
		optNilStr = append(optNilStr, "Host")
		fallthrough
	case len(opt.Port) < 2:
		optNilStr = append(optNilStr, "Port")
		fallthrough
	case len(opt.Database) < 2:
		optNilStr = append(optNilStr, "Database")
		fallthrough
	case len(opt.Collection) < 2:
		optNilStr = append(optNilStr, "Collection")
	}

	if len(optNilStr) > 0 {
		fmt.Println("缺少参数:", strings.Join(optNilStr, ","))
	}

	fmt.Println("参数通过")

	var DB *DB

	return DB
}
