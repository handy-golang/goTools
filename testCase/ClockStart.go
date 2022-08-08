package testCase

import (
	"fmt"
	"time"

	"github.com/EasyGolang/goTools/mClock"
)

func ClockStart() {
	mClock.New(mClock.OptType{
		Func: func() {
			fmt.Println(time.Now())
		},
		Spec: "0/1 * * * * ?",
	})
}
