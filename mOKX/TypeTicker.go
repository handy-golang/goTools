package mOKX

// 币安原始榜单数据
type TypeBinanceTicker struct {
	Symbol             string `bson:"symbol"`
	InstID             string `bson:"InstID"`
	PriceChange        string `bson:"priceChange"`
	PriceChangePercent string `bson:"priceChangePercent"`
	WeightedAvgPrice   string `bson:"weightedAvgPrice"`
	PrevClosePrice     string `bson:"prevClosePrice"`
	LastPrice          string `bson:"lastPrice"`
	LastQty            string `bson:"lastQty"`
	BidPrice           string `bson:"bidPrice"`
	BidQty             string `bson:"bidQty"`
	AskPrice           string `bson:"askPrice"`
	AskQty             string `bson:"askQty"`
	OpenPrice          string `bson:"openPrice"`
	HighPrice          string `bson:"highPrice"`
	LowPrice           string `bson:"lowPrice"`
	Volume             string `bson:"volume"`
	QuoteVolume        string `bson:"quoteVolume"`
	OpenTime           int64  `bson:"openTime"`
	CloseTime          int64  `bson:"closeTime"`
	FirstID            int    `bson:"firstId"`
	LastID             int    `bson:"lastId"`
	Count              int    `bson:"count"`
}

// OKX 原始榜单数据
type TypeOKXTicker struct {
	InstType  string `bson:"instType"`
	InstID    string `bson:"instId"`
	Last      string `bson:"last"`
	LastSz    string `bson:"lastSz"`
	AskPx     string `bson:"askPx"`
	AskSz     string `bson:"askSz"`
	BidPx     string `bson:"bidPx"`
	BidSz     string `bson:"bidSz"`
	Open24H   string `bson:"open24h"`
	High24H   string `bson:"high24h"`
	Low24H    string `bson:"low24h"`
	VolCcy24H string `bson:"volCcy24h"`
	Vol24H    string `bson:"vol24h"`
	Ts        string `bson:"ts"`
	SodUtc0   string `bson:"sodUtc0"`
	SodUtc8   string `bson:"sodUtc8"`
}

// 综合榜单数据
type TypeTicker struct {
	InstID         string   `bson:"InstID"` // 产品ID
	Symbol         string   `bson:"symbol"`
	CcyName        string   `bson:"CcyName"`        // 币种名称
	Last           string   `bson:"Last"`           // 最新成交价
	Open24H        string   `bson:"Open24H"`        // 24小时开盘价
	High24H        string   `bson:"High24H"`        // 最高价
	Low24H         string   `bson:"Low24H"`         // 最低价
	OKXVol24H      string   `bson:"OKXVol24H"`      // OKX 24小时成交量 USDT 数量
	BinanceVol24H  string   `bson:"BinanceVol24H"`  // 24 小时成交 USDT 数量
	U_R24          string   `bson:"U_R24"`          // 涨幅 = (最新价-开盘价)/开盘价 =
	Volume         string   `bson:"Volume"`         // 成交量总和
	OkxVolRose     string   `bson:"OkxVolRose"`     // 欧意占比总交易量
	BinanceVolRose string   `bson:"BinanceVolRose"` // 币安占比总交易量
	Ts             int64    `bson:"Ts"`
	TimeUnix       int64    `bson:"TimeUnix"`
	TimeStr        string   `bson:"TimeStr"`
	SPOT           TypeInst `bson:"	SPOT"`
	SWAP           TypeInst `bson:"	SWAP"`
}

// 基于 TickerList  的市场分析
type TypeWholeTickerAnaly struct {
	DiffHour      int            `bson:"DiffHour"`     // 总时长	 分析的切片时长
	UPIndex       string         `bson:"UPIndex"`      // 上涨指数  上涨个数 / 下跌个数   * 100%
	UDAvg         string         `bson:"UDAvg"`        // 综合涨幅均值  上涨平均值 + 下跌平均值
	UPLe          int            `bson:"UPLe"`         // 上涨趋势  上涨指数 大于 50%
	UDLe          int            `bson:"UDLe"`         // 上涨强度  综合涨幅均值 大于 0
	DirIndex      int            `bson:"DirIndex"`     // 当前市场情况  -1 下跌   0 震荡   1 上涨
	MaxUP         AnalySliceType `bson:"MaxUP"`        // 最大涨幅币种
	MaxDown       AnalySliceType `bson:"MaxDown"`      // 最大跌幅币种
	StartTimeStr  string         `bson:"StartTimeStr"` // 开始时间
	StartTimeUnix int64          `bson:"StartTimeUnix"`
	EndTimeStr    string         `bson:"EndTimeStr"` // 结束时间
	EndTimeUnix   int64          `bson:"EndTimeUnix"`
}
