package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/mTime"
)

func TimeTest() {
	fmt.Println(mTime.UnixFormat(mTime.GetUnixInt64()))
}
