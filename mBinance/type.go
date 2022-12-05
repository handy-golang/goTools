package mBinance

type PositionType struct {
	PositionID    string `bson:"PositionID"`
	InstID        string `bson:"InstID"`
	Dir           int    `bson:"Dir"`
	Profit        string `bson:"Profit"` // 持仓数量
	CreateTime    int64  `bson:"CreateTime"`
	CreateTimeStr string `bson:"CreateTimeStr"`
}
