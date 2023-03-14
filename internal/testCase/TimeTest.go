package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/mTime"
)

func TimeTest() {
	// fmt.Println(mTime.UnixFormat(mTime.GetUnixInt64()))

	TimeStr := "2023-05-06 18:56:43"

	unix := mTime.TimeParse(mTime.LaySP_ss, TimeStr)

	fmt.Println(unix, mTime.UnixFormat(unix))
}
