package mMongo

/*
https://github.com/mongodb/mongo-go-driver/
*/

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Start() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	Url := "mongodb://mo7:asdasd55555@mo7.cc:17017/Hunter"
	// 建立连接
	o := options.Client().ApplyURI(Url)
	db, err := mongo.Connect(ctx, o)
	if err != nil {
		errStr := fmt.Errorf("连接失败: %+v", err)
		fmt.Println(errStr)
		return
	}

	err = db.Ping(ctx, readpref.Primary())
	if err != nil {
		errStr := fmt.Errorf("认证失败: %+v", err)
		fmt.Println(errStr)
		return
	}

	table := db.Database("Hunter").Collection("HotList")
	// Client := base.Client()

	fmt.Println("我们继续", table)
}
