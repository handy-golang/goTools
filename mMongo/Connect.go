package mMongo

import (
	"context"
	"fmt"
	"time"

	"github.com/EasyGolang/goTools/mStr"
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
		Name string `bson:"name"`
		Age  int    `bson:"age"`
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
		Name string `bson:"name"`
		Age  int    `bson:"age"`
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

	https://blog.csdn.net/weixin_44738411/article/details/106347939

	https://blog.csdn.net/weixin_44738411/article/details/104276995

*/

// 连接数据库
func (info *DB) Connect() *DB {
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

	info.Event("Connected", "连接成功")

	info.db = info.Client.Database(info.dbName)

	info.Event("Database", info.dbName)

	return info
}

func (info *DB) Ping() error {
	err := info.Client.Ping(info.Ctx, readpref.Primary())
	return err
}

func (info *DB) Close() {
	info.cancel()
	info.Event("Close", mStr.ToStr(info))
}

func (info *DB) Collection(tableName string) *DB {
	info.Table = info.db.Collection(tableName)
	info.Event("Collection", tableName)
	return info
}
