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
)

// 将13位毫秒的时间戳字符串转为 时间对象，如果不是13 位则返回当前时间
func MsToTime(ms string, diff string) time.Time {
	msStr := mCount.Add(ms, diff)

	if len(ms) == 13 {
		msInt, _ := strconv.ParseInt(msStr, 10, 64)
		tm := time.Unix(0, msInt*int64(time.Millisecond))
		return tm
	}
	return time.Now()
}

// 获取 13 位毫秒时间戳
func GetUnix() string {
	unix := time.Now().UnixNano() / 1e6
	str := strconv.FormatInt(unix, 10)

	return str
}

// 2022-02-23T13:39:24.630Z
func IsoTime() string {
	utcTime := time.Now().UTC()
	iso := utcTime.String()
	isoBytes := []byte(iso)
	iso = string(isoBytes[:10]) + "T" + string(isoBytes[11:23]) + "Z"
	return iso
}

func EpochTime() string {
	millisecond := time.Now().UnixNano() / 1000000
	epoch := strconv.Itoa(int(millisecond))
	epochBytes := []byte(epoch)
	epoch = string(epochBytes[:10])

	return epoch
}
