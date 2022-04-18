package main

import (
	"fmt"
	"time"

	"github.com/meichangliang/goTools/mCycle"
	"github.com/meichangliang/goTools/mTime"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	in := 0

	Ticker := func() {
		Unix := mTime.GetUnix()
		time := mTime.MsToTime(Unix, "0")
		in++
		fmt.Println("Ticker", time, in)
	}

	t := mCycle.New(mCycle.CycleParam{
		Func:      Ticker,
		SleepTime: time.Second,
	})

	t.Start()

	time.Sleep(time.Second * 10)

	t.End()

	time.Sleep(time.Hour)

	fmt.Println(" =========   END   ========= ")
}
