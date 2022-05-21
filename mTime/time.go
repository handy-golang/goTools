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

// 将时间戳转为 时间对象，如果不正确 则返回当前时间对象
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

// 格式化时间戳
func UnixFormat(ms string) string {
	timeMs := ms
	if len(ms) < 1 {
		timeMs = GetUnix()
	}
	T := MsToTime(timeMs, "0")
	timeStr := T.Format("2006-01-02T15:04:05")

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

func IsoTime(isUTC bool) string {
	utcTime := time.Now()
	if isUTC {
		utcTime = time.Now().UTC()
	}

	iso := utcTime.String()
	isoBytes := []byte(iso)
	iso = string(isoBytes[:10]) + "T" + string(isoBytes[11:23]) + "Z"
	return iso
}

func RFCTime(lType bool) string {
	TimeData := time.Now().Format(time.RFC3339Nano)
	if lType {
		TimeData = time.Now().Format(time.RFC3339)
	}
	return TimeData
}
