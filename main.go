package main

import (
	_ "embed"
	"fmt"

	"github.com/EasyGolang/goTools/global"
	"github.com/EasyGolang/goTools/global/config"
	"github.com/EasyGolang/goTools/mJson"
	"github.com/EasyGolang/goTools/mMongo"
	"github.com/EasyGolang/goTools/mOKX"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
)

//go:embed package.json
var AppPackage []byte

func main() {
	jsoniter.Unmarshal(AppPackage, &config.AppInfo)
	global.Start()

	fmt.Println(" =========  START  ========= ")

	// testCase.ClockStart()

	// testCase.GetSPOT()

	// testCase.GetKdata("EOS-USDT")

	// testCase.OKXFetch()
	// testCase.OKXWss()

	// testCase.CountTest()
	// testCase.StrFuzzy()

	// testCase.FileTest()

	// testCase.YaSuoDir()
	// OrganizeDatabase()
	fmt.Println(" =========   END   ========= ")
}

type MarketTickerTable struct {
	List           []mOKX.TypeTicker                `bson:"List"`        // 成交量排序列表
	ListU_R24      []mOKX.TypeTicker                `bson:"ListU_R24"`   // 涨跌幅排序列表
	AnalyWhole     []mOKX.TypeWholeTickerAnaly      `bson:"AnalyWhole"`  // 大盘分析结果
	AnalySingle    map[string][]mOKX.AnalySliceType `bson:"AnalySingle"` // 单个币种分析结果
	Unit           string                           `bson:"Unit"`
	WholeDir       int                              `bson:"WholeDir"`
	TimeUnix       int64                            `bson:"TimeUnix"`
	Time           string                           `bson:"Time"`
	CreateTimeUnix int64                            `bson:"CreateTimeUnix"`
	CreateTime     string                           `bson:"CreateTime"`
}

type MarketTickerAPI struct {
	Unit            string `bson:"Unit"`
	WholeDir        int    `bson:"WholeDir"`
	TimeUnix        int64  `bson:"TimeUnix"`
	Time            string `bson:"Time"`
	CreateTimeUnix  int64  `bson:"CreateTimeUnix"`
	CreateTime      string `bson:"CreateTime"`
	MaxUP           string `bson:"MaxUP"` // 最大涨幅币种
	MaxUP_RosePer   string `bson:"MaxUP_RosePer"`
	MaxDown         string `bson:"MaxDown"` // 最大跌幅币种
	MaxDown_RosePer string `bson:"MaxDown_RosePer"`
}

func OrganizeDatabase() {
	db := mMongo.New(mMongo.Opt{
		UserName: config.AppEnv.MongoUserName,
		Password: config.AppEnv.MongoPassword,
		Address:  config.AppEnv.MongoAddress,
		DBName:   "AITrade",
	}).Connect().Collection("MarketTicker")
	defer db.Close()
	cur, err := db.Table.Find(db.Ctx, bson.D{})
	if err != nil {
		fmt.Println("数据库错误")
		return
	}
	var MarketTickerList []MarketTickerAPI
	Key := 0

	for cur.Next(db.Ctx) {
		var result MarketTickerTable
		cur.Decode(&result)

		var MarketTicker MarketTickerAPI
		jsoniter.Unmarshal(mJson.ToJson(result), &MarketTicker)
		MarketTicker.MaxUP = result.AnalyWhole[0].MaxUP.CcyName
		MarketTicker.MaxUP_RosePer = result.AnalyWhole[0].MaxUP.RosePer
		MarketTicker.MaxDown = result.AnalyWhole[0].MaxDown.CcyName
		MarketTicker.MaxDown_RosePer = result.AnalyWhole[0].MaxDown.RosePer
		MarketTickerList = append(MarketTickerList, MarketTicker)

		Val := MarketTicker

		if Key > 1 {
			Pre := MarketTickerList[Key-1]
			fmt.Println(Key, Val.TimeUnix-Pre.TimeUnix)
		}

		Key++
	}
}
