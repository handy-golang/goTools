package mOKX

import "time"

type CandleDataType [7]string

type Kd struct {
	InstID   string    `json:"InstID"`   // 持仓币种
	TimeUnix int64     `json:"TimeUnix"` // 毫秒时间戳
	Time     time.Time `json:"Time"`     // 时间
	O        string    `json:"O"`        // 开盘
	H        string    `json:"H"`        // 最高
	L        string    `json:"L"`        // 最低
	C        string    `json:"C"`        // 收盘价格
	Vol      string    `json:"Vol"`      // 交易货币的数量
	VolCcy   string    `json:"VolCcy"`   // 计价货币数量
	Type     string    `json:"Type"`     // 数据类型
	Dir      int       `json:"Dir"`      // 方向 (收盘-开盘) ，1：涨 & -1：跌 & 0：横盘
	Center   string    `json:"Center"`   // 实体中心价 (开盘+收盘+最高+最低) / 4
	HLPer    string    `json:"HLPer"`    // (最高-最低)/最低 * 100%
	U_shade  string    `json:"U_shade"`  // 上影线
	D_shade  string    `json:"D_shade"`  // 下影线
	RosePer  string    `json:"RosePer"`  // 涨幅 当前收盘价 - 上一位收盘价 * 100%
	C_dir    int       `json:"C_dir"`    // 中心点方向 (当前中心点-前中心点) 1：涨 & -1：跌 & 0：横盘
	H_idx    int       `json:"H_idx"`    // 最高价排序序号
	L_idx    int       `json:"L_idx"`    // 按照最低价排序
	C_idx    int       `json:"C_idx"`    // 按照收盘价排序
}

// 基于 K线数据分析结果
type AnalyseSliceType struct {
	StartTime     time.Time `json:"StartTime"` // 开始时间
	StartTimeUnix int64     `json:"StartTimeUnix"`
	EndTime       time.Time `json:"EndTime"` // 结束时间
	EndTimeUnix   int64     `json:"EndTimeUnix"`
	DiffHour      int       `json:"DiffHour"` // 总时长
	Len           int       `json:"Len"`      // 数据的总长度
	Volume        string    `json:"Volume"`   // 成交量总和
	RosePer       string    `json:"RosePer"`  // 涨幅
}

type AnalyseSingleType struct {
	InstID        string    `json:"InstID"`    // InstID
	StartTime     time.Time `json:"StartTime"` // 开始时间
	StartTimeUnix int64     `json:"StartTimeUnix"`
	EndTime       time.Time `json:"EndTime"` // 结束时间
	EndTimeUnix   int64     `json:"EndTimeUnix"`
	DiffHour      int64     `json:"DiffHour"` // 总时长
}
