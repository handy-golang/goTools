package binance

type PositionType struct {
	InstID        string `bson:"InstID"`
	Dir           int    `bson:"Dir"`
	Profit        string `bson:"Profit"` // 未实现盈亏
	CreateTime    int64  `bson:"CreateTime"`
	CreateTimeStr string `bson:"CreateTimeStr"`
	UpdateTime    int64  `bson:"UpdateTime"`
	UpdateTimeStr string `bson:"UpdateTimeStr"`
}
