package mTime

/*

//13 位时间戳字符串，第二位参数可加减时间戳，
// 注：如果时间不对则返回当前时间
str := "1489582166978"
myTime := MsToTime(str, "-1000")


*/

import (
	"strconv"
	"time"

	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mStr"
)

/*
将毫秒时间戳转为 时间对象，如果不正确 则返回当前时间对象
ms =  string | int64 毫秒数   diff = "-988"
*/
func MsToTime(ms any, diff string) time.Time {
	msToStr := mStr.ToStr(ms)

	msStr := mCount.Add(msToStr, diff)

	msInt, err := strconv.ParseInt(msStr, 10, 64)
	if err != nil {
		return time.Now()
	}
	tm := time.Unix(0, msInt*int64(time.Millisecond))
	return tm
}

// ms=string | int64 毫秒数  return = 2006-01-02T15:04:05
func UnixFormat(ms any) string {
	timeMs := mStr.ToStr(ms)
	if len(timeMs) < 1 {
		timeMs = GetUnix()
	}
	T := MsToTime(timeMs, "0")
	timeStr := T.Format(Lay_ss)
	return timeStr
}

// 获取 13 位毫秒时间戳
func GetUnix() string {
	unix := time.Now().UnixNano() / 1e6
	str := strconv.FormatInt(unix, 10)
	return str
}

func GetUnixInt64() int64 {
	return time.Now().UnixNano() / 1e6
}

// 将时间对象转为毫秒
func ToUnixMsec(ms time.Time) int64 {
	return ms.UnixNano() / 1e6
}

// 2020-12-03 转为 时间戳  mTime.TimeParse(mTime.LaySP_ss, "2023-05-06 18:56:43")
func TimeParse(layout, val string) (resData int64) {
	formatTime, err := time.ParseInLocation(layout, val, time.Local)
	if err != nil {
		return
	}
	resData = ToUnixMsec(formatTime)
	return
}
