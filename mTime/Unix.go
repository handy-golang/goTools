package mTime

type UnixTimeType struct {
	Seconds string `bson:"Seconds"`
	Minute  string `bson:"Minute"`
	Hour    string `bson:"Hour"`
	Day     string `bson:"Day"`
}

// 毫秒数
var UnixTime = UnixTimeType{
	Seconds: "1000",     // 1 秒钟 毫秒数
	Minute:  "60000",    // 1 分钟 毫秒数
	Hour:    "3600000",  // 1 小时 毫秒数
	Day:     "86400000", // 1 天   毫秒数
}

type UnixTimeInt64Type struct {
	Seconds int64 `bson:"Seconds"`
	Minute  int64 `bson:"Minute"`
	Hour    int64 `bson:"Hour"`
	Day     int64 `bson:"Day"`
}

var UnixTimeInt64 = UnixTimeInt64Type{
	Seconds: 1000,     // 1 秒钟 毫秒数
	Minute:  60000,    // 1 分钟 毫秒数
	Hour:    3600000,  // 1 小时 毫秒数
	Day:     86400000, // 1 天   毫秒数
}

// 时间解析字符串模板

var (
	LaySP_ss = "2006-01-02 15:04:05"
	LaySP_s  = "2006-1-2 15:4:5"

	Lay_ss = "2006-01-02T15:04:05"
	Lay_s  = "2006-1-2T15:4:5"
	Lay_DD = "2006-01-02"
	Lay_D  = "2006-1-2"
	Lay_MM = "2006-01"
	Lay_M  = "2006-1"
	Lay_Y  = "2006"
)
