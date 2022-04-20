package mMongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// 连接数据库
func (info *DBInfo) Connect() *DBInfo {
	info.Ctx, info.cancel = context.WithTimeout(
		context.Background(),
		time.Duration(info.Timeout)*time.Second,
	)

	o := options.Client().ApplyURI(info.URI)
	Client, err := mongo.Connect(info.Ctx, o)
	if err != nil {
		errStr := fmt.Sprintf("连接失败: %+v", err)
		info.Event("ConnectErr", errStr)
		return info
	}

	err = Client.Ping(info.Ctx, readpref.Primary())
	if err != nil {
		errStr := fmt.Sprintf("验证失败: %+v", err)
		info.Event("PingErr", errStr)
		return info
	}

	info.Client = Client

	info.db = info.Client.Database(info.dbName)

	return info
}

func (info *DBInfo) Ping() error {
	err := info.Client.Ping(info.Ctx, readpref.Primary())
	return err
}

func (info *DBInfo) Close() {
	info.cancel()
}

func (info *DBInfo) Collection(tableName string) *DBInfo {
	info.Table = info.db.Collection(tableName)
	return info
}
