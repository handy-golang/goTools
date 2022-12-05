package mBinance

type PositionType struct {
	InstID        string `bson:"InstID"`
	Dir           int    `bson:"Dir"`
	Profit        string `bson:"Profit"` // 持仓数量
	CreateTime    int64  `bson:"CreateTime"`
	CreateTimeStr string `bson:"CreateTimeStr"`
}
