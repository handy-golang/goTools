package mTime

type UnixTimeType struct {
	Seconds string `json:"Seconds"`
	Minute  string `json:"Minute"`
	Hour    string `json:"Hour"`
	Day     string `json:"Day"`
}

// 毫秒数
var UnixTime = UnixTimeType{
	Seconds: "1000",     // 1 秒钟 毫秒数
	Minute:  "60000",    // 1 分钟 毫秒数
	Hour:    "3600000",  // 1 小时 毫秒数
	Day:     "86400000", // 1 天   毫秒数
}

type UnixTimeInt64Type struct {
	Seconds int64 `json:"Seconds"`
	Minute  int64 `json:"Minute"`
	Hour    int64 `json:"Hour"`
	Day     int64 `json:"Day"`
}

var UnixTimeInt64 = UnixTimeInt64Type{
	Seconds: 1000,     // 1 秒钟 毫秒数
	Minute:  60000,    // 1 分钟 毫秒数
	Hour:    3600000,  // 1 小时 毫秒数
	Day:     86400000, // 1 天   毫秒数
}
