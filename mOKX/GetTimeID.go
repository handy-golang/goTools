package mOKX

import "github.com/EasyGolang/goTools/mTime"

func GetTimeID(TimeUnix int64) string {
	T := mTime.MsToTime(TimeUnix, "0")
	timeStr := T.Format("2006-01-02T15:04")
	return timeStr
}
