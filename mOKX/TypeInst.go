package mOKX

type TypeInst struct {
	Alias        string `bson:"Alias"`
	BaseCcy      string `bson:"BaseCcy"`
	Category     string `bson:"Category"`
	CtMult       string `bson:"CtMult"`
	CtType       string `bson:"CtType"`
	CtVal        string `bson:"CtVal"`
	CtValCcy     string `bson:"CtValCcy"`
	ExpTime      string `bson:"ExpTime"`
	InstID       string `bson:"InstID"`
	InstType     string `bson:"InstType"`
	Lever        string `bson:"Lever"`
	ListTime     string `bson:"ListTime"`
	LotSz        string `bson:"LotSz"`
	MaxIcebergSz string `bson:"MaxIcebergSz"`
	MaxLmtSz     string `bson:"MaxLmtSz"`
	MaxMktSz     string `bson:"MaxMktSz"`
	MaxStopSz    string `bson:"MaxStopSz"`
	MaxTriggerSz string `bson:"MaxTriggerSz"`
	MaxTwapSz    string `bson:"MaxTwapSz"`
	MinSz        string `bson:"MinSz"`
	OptType      string `bson:"OptType"`
	QuoteCcy     string `bson:"QuoteCcy"`
	SettleCcy    string `bson:"SettleCcy"`
	State        string `bson:"State"`
	Stk          string `bson:"Stk"`
	TickSz       string `bson:"TickSz"`
	Uly          string `bson:"Uly"`
}

type TypeReq struct {
	Code string `bson:"Code"`
	Data any    `bson:"Data"`
	Msg  string `bson:"Msg"`
}
