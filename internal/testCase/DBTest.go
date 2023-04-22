package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	"github.com/EasyGolang/goTools/mStruct"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 连通篇
func StartDBRun() {
	db, err := mMongo.New(mMongo.Opt{
		UserName: "123",
		Password: "123",
		Address:  "xx.xx.xxx:xxx",
		Timeout:  100, // 秒
		DBName:   "xxxx",
	}).Connect()
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer db.Close()
	db = db.Collection("aaaaa")

	CountDocuments(db)

	// Insert(db)
	// InsertMany(db)

	// UpdateOne(db)

	// UpdateMany(db)

	// Drop(db)

	// Find(db)

	// fmt.Println(db)

	// FindOne(db)
}

// 查询文档总个数
func CountDocuments(db *mMongo.DB) {
	FK := bson.D{}
	CountOpt := options.CountOptions{}
	total, err := db.Table.CountDocuments(db.Ctx, FK, &CountOpt)
	fmt.Println(total, err)
}

// 插入文档单个
func Insert(db *mMongo.DB) {
	var InsertData mOKX.TypeKd
	InsertData.TimeUnix = 1

	// 约等于没有设置项
	InsertOpt := options.InsertOne()
	db.Table.InsertOne(db.Ctx, InsertData, InsertOpt)
}

// 批量插入文档
func InsertMany(db *mMongo.DB) {
	var dbInsertList []any

	for i := 0; i < 30; i++ {
		var InsertData mOKX.TypeKd
		InsertData.TimeUnix = int64(i)

		dbInsertList = append(dbInsertList, InsertData)
	}

	// 约等于没有设置项
	InsertOpt := options.InsertMany()
	db.Table.InsertMany(db.Ctx, dbInsertList, InsertOpt)
}

// UpDate Document 可以完美取代 InsertOne
func UpdateOne(db *mMongo.DB) {
	var InsertData mOKX.TypeKd
	InsertData.TimeUnix = 1
	InsertData.TimeStr = "234"
	InsertData.TimeUnix++

	FK := bson.D{{
		Key:   "TimeUnix",
		Value: InsertData.TimeUnix,
	}}
	UK := bson.D{}
	mStruct.Traverse(InsertData, func(key string, val any) {
		UK = append(UK, bson.E{
			Key: "$set",
			Value: bson.D{
				{
					Key:   key,
					Value: val,
				},
			},
		})
	})

	// 约等于没有设置项
	upOpt := options.Update()
	upOpt.SetUpsert(true) // 如果匹配不到则插入
	db.Table.UpdateOne(db.Ctx, FK, UK, upOpt)
}

// 修改多个
func UpdateMany(db *mMongo.DB) {
	var InsertData mOKX.TypeKd
	InsertData.TimeUnix = 1
	InsertData.TimeStr = "234"
	InsertData.TimeUnix++

	FK := bson.D{{
		Key:   "TimeUnix",
		Value: InsertData.TimeUnix,
	}}
	UK := bson.D{}
	mStruct.Traverse(InsertData, func(key string, val any) {
		UK = append(UK, bson.E{
			Key: "$set",
			Value: bson.D{
				{
					Key:   key,
					Value: val,
				},
			},
		})
	})

	// 约等于没有设置项
	upOpt := options.Update()
	upOpt.SetUpsert(true) // 如果匹配不到则插入
	db.Table.UpdateMany(db.Ctx, FK, UK, upOpt)
}

// 删除数据
func Drop(db *mMongo.DB) {
	// 约等于没有设置项
	// upOpt := options.Update()
	// upOpt.SetUpsert(true) // 如果匹配不到则插入
	// db.Table.Drop(db.Ctx) // 删除集合

	FK := bson.D{{
		Key:   "TimeUnix",
		Value: 2,
	}}

	db.Table.DeleteOne(db.Ctx, FK) // 删除数据
}

// 查询多个
func Find(db *mMongo.DB) {
	FK := bson.D{}

	findOpt := options.Find()
	findOpt.SetSort(map[string]int{
		"TimeUnix": 1,
	})
	// findOpt.SetSkip(1) // 跳过几页
	// findOpt.SetLimit(2) // 设置最大数量
	findOpt.SetAllowDiskUse(true) // 查询数据必备

	cur, err := db.Table.Find(db.Ctx, FK, findOpt)

	for cur.Next(db.Ctx) {
		var curData map[string]any
		cur.Decode(&curData)
		fmt.Println(curData["TimeUnix"])
	}

	fmt.Println(err)
}

func FindOne(db *mMongo.DB) {
	FK := bson.D{{
		Key:   "TimeUnix",
		Value: 6,
	}}

	findOpt := options.FindOne()
	findOpt.SetSort(map[string]int{
		"TimeUnix": -1,
	})
	// findOpt.SetSkip(1) // 跳过几页
	// findOpt.SetLimit(2) // 设置最大数量
	var curData map[string]any
	db.Table.FindOne(db.Ctx, FK, findOpt).Decode(&curData)

	fmt.Println(curData["TimeUnix"])
}
