package mOKX

type TypeInst struct {
	Alias        string `bson:"alias"`
	BaseCcy      string `bson:"baseCcy"`
	Category     string `bson:"category"`
	CtMult       string `bson:"ctMult"`
	CtType       string `bson:"ctType"`
	CtVal        string `bson:"ctVal"`
	CtValCcy     string `bson:"ctValCcy"`
	ExpTime      string `bson:"expTime"`
	InstID       string `bson:"instId"`
	InstType     string `bson:"instType"`
	Lever        string `bson:"lever"`
	ListTime     string `bson:"listTime"`
	LotSz        string `bson:"lotSz"`
	MaxIcebergSz string `bson:"maxIcebergSz"`
	MaxLmtSz     string `bson:"maxLmtSz"`
	MaxMktSz     string `bson:"maxMktSz"`
	MaxStopSz    string `bson:"maxStopSz"`
	MaxTriggerSz string `bson:"maxTriggerSz"`
	MaxTwapSz    string `bson:"maxTwapSz"`
	MinSz        string `bson:"minSz"`
	OptType      string `bson:"optType"`
	QuoteCcy     string `bson:"quoteCcy"`
	SettleCcy    string `bson:"settleCcy"`
	State        string `bson:"state"`
	Stk          string `bson:"stk"`
	TickSz       string `bson:"tickSz"`
	Uly          string `bson:"uly"`
}

type TypeReq struct {
	Code string `bson:"code"`
	Data any    `bson:"data"`
	Msg  string `bson:"msg"`
}
