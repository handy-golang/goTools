package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/mTime"
)

func TimeTest() {
	// fmt.Println(mTime.UnixFormat(mTime.GetUnixInt64()))

	// TimeStr := "2023-05-06 18:56:43"

	// unix := mTime.TimeParse(mTime.LaySP_ss, TimeStr)

	// fmt.Println(unix, mTime.UnixFormat(unix))

	// 测试，将 毫秒 时间 转为   x 时 x 分钟 x 秒
	unix := mTime.UnixTimeInt64.Hour*5 +
		mTime.UnixTimeInt64.Minute*3 +
		mTime.UnixTimeInt64.Seconds*8

	hms := mTime.UnixTo_hh_mm_ss(unix)

	fmt.Println(hms)
}
