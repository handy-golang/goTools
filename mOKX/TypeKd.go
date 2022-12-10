package mOKX

// K 线 数据
type TypeKd struct {
	InstID   string `bson:"InstID"`   // 持仓币种
	TimeUnix int64  `bson:"TimeUnix"` // 毫秒时间戳
	TimeStr  string `bson:"TimeStr"`  // 时间的字符串形式
	O        string `bson:"O"`        // 开盘
	H        string `bson:"H"`        // 最高
	L        string `bson:"L"`        // 最低
	C        string `bson:"C"`        // 收盘价格
	CBas     string `bson:"CBas"`     // 实体中心价 (收盘+最高+最低) / 3
	Vol      string `bson:"Vol"`      // 成交量( BTC 的数量 )
	Dir      int    `bson:"Dir"`      // 方向 (收盘-开盘) ，1：涨 & -1：跌 & 0：横盘
	HLPer    string `bson:"HLPer"`    // 振幅 (最高-最低)/最低 * 100%
	U_shade  string `bson:"U_shade"`  // 上影线
	D_shade  string `bson:"D_shade"`  // 下影线
	RosePer  string `bson:"RosePer"`  // 涨幅 当前收盘价 - 上一位收盘价 * 100%
	C_dir    int    `bson:"C_dir"`    // 实体中心价 (当前中心点-前中心点) 1：涨 & -1：跌 & 0：横盘
}

// 基于 K线数据分析结果
type AnalySliceType struct {
	InstID        string `bson:"InstID"`       // InstID
	CcyName       string `bson:"CcyName"`      // 币种名称
	StartTimeStr  string `bson:"StartTimeStr"` // 开始时间
	StartTimeUnix int64  `bson:"StartTimeUnix"`
	EndTimeStr    string `bson:"EndTimeStr"` // 结束时间
	EndTimeUnix   int64  `bson:"EndTimeUnix"`
	DiffHour      int    `bson:"DiffHour"`   // 总时长
	Len           int    `bson:"Len"`        // 数据的总长度
	Volume        string `bson:"Volume"`     // 成交量总和
	VolumeAvg     string `bson:"VolumeAvg"`  // 平均 小时 成交量
	RosePer       string `bson:"RosePer"`    // 涨幅
	H             string `bson:"H"`          // 最高价
	L             string `bson:"L"`          // 最低价
	U_shadeAvg    string `bson:"U_shadeAvg"` // 上影线平均长度
	D_shadeAvg    string `bson:"D_shadeAvg"` // 下影线平均长度
	HLPerMax      string `bson:"HLPerMax"`   // 最高振幅
	HLPerAvg      string `bson:"HLPerAvg"`   // 平均振幅
}

type AnalySingleType struct {
	InstID        string `bson:"InstID"`       // InstID
	StartTimeStr  string `bson:"StartTimeStr"` // 开始时间
	StartTimeUnix int64  `bson:"StartTimeUnix"`
	EndTimeStr    string `bson:"EndTimeStr"` // 结束时间
	EndTimeUnix   int64  `bson:"EndTimeUnix"`
	DiffHour      int64  `bson:"DiffHour"` // 总时长
}
