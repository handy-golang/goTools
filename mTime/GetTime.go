package mTime

type GetTimeReturnType struct {
	TimeUnix int64  `bson:"TimeUnix"`
	TimeStr  string `bson:"TimeStr"`
}

func GetTime() (resData GetTimeReturnType) {
	resData.TimeUnix = GetUnixInt64()
	resData.TimeStr = UnixFormat(resData.TimeUnix)
	return
}

// ms 为 毫秒 时间戳
func TimeGet(ms any) (resData GetTimeReturnType) {
	myTime := MsToTime(ms, "0")
	resData.TimeUnix = ToUnixMsec(myTime)
	resData.TimeStr = UnixFormat(resData.TimeUnix)
	return
}
