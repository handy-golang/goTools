package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/mTime"
)

func TimeTest() {
	fmt.Println(mTime.IsoTime())
	fmt.Println(mTime.RFCTime(false))
	fmt.Println(mTime.RFCTime(true))
	fmt.Println(mTime.EpochTime())
	fmt.Println(mTime.GetUnixInt64())

	fmt.Println(mTime.UnixFormat(mTime.GetUnixInt64()))
}
