package mMongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*

	db := mMongo.New(mMongo.Opt{
		UserName: "aadd",
		Password: "123456",
		Host:     "127.235.565.415",
		Port:     "27017",
		DBName:   "Hunter",
	}).Connect().Collection("HotList")

	https://mongoing.com/archives/27257

*/

/*

	插入操作
	type body struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var jk []any
	for i := 0; i < 10; i++ {
		b := body{}
		b.Name = "iii" + mStr.ToStr(i)
		b.Age = 1 + i
		jk = append(jk, b)
	}
	insertManyResult, err := db.Table.InsertMany(db.Ctx, jk)
	if err != nil {
		fmt.Println("插入出错")
	}

*/

/*
	// 查询操作
	var result struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	filter := bson.D{
		{"name", "iii5"},
	}

	err := db.Table.FindOne(db.Ctx, filter).Decode(&result)
	if err != nil {
		fmt.Println("查询错误", err)
	}

	fmt.Println("查询结果", string(mJson.ToJson(result)))

*/

/*
// 修改
	filter := bson.D{
		{"name", "iii6"},
	}

	update := bson.D{
		{"$inc", bson.D{
			{"age", 548},
		}},
	}

	updateResult, err := db.Table.UpdateOne(db.Ctx, filter, update)
	if err != nil {
		fmt.Println("修改错误", err)
	}

	fmt.Println("修改结果", mStr.ToStr(updateResult))


*/

/*

	// 删除
	filter := bson.D{
		{"name", "iii9"},
	}
	deleteResult, err := db.Table.DeleteOne(db.Ctx, filter)
	if err != nil {
		fmt.Println("删除失败", err)
	}

*/

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
