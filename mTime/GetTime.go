package mTime

import (
	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mStr"
)

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

// 将毫秒时间戳转换为 x 时 x 分 x 秒

type HMSType struct {
	HH  string
	MM  string
	SS  string
	HMS string
}

func UnixTo_hh_mm_ss(ms any) HMSType {
	msToStr := mStr.ToStr(ms)

	H := 0
	M := 0
	S := "0"

	// 计算小时
	in_h := mCount.Div(msToStr, UnixTime.Hour)
	in_m := "0"
	in_s := "0"

	H = mCount.ToInt(in_h)

	h_dec := mCount.GetDecimal(in_h)
	if h_dec > 0 {
		// 计算分钟
		// 减去小时
		H_unix := UnixTimeInt64.Hour * int64(H)
		in_m_unix := mCount.Sub(msToStr, mStr.ToStr(H_unix))
		in_m = mCount.Div(in_m_unix, UnixTime.Minute)
	}
	M = mCount.ToInt(in_m)

	m_dec := mCount.GetDecimal(in_m)
	if m_dec > 0 {
		// 计算秒  减去 分钟 和 小时
		HM_unix := UnixTimeInt64.Hour*int64(H) + UnixTimeInt64.Minute*int64(M)
		in_s_unix := mCount.Sub(msToStr, mStr.ToStr(HM_unix))
		in_s = mCount.Div(in_s_unix, UnixTime.Seconds)
	}
	S = mCount.CentRound(in_s, 1)

	ReturnS := ""
	ReturnS = mStr.Join(ReturnS, H, "时")
	ReturnS = mStr.Join(ReturnS, M, "分")
	ReturnS = mStr.Join(ReturnS, S, "秒")

	return HMSType{
		HH:  mStr.ToStr(H),
		MM:  mStr.ToStr(M),
		SS:  S,
		HMS: ReturnS,
	}
}
