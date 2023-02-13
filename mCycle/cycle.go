package mCycle

import (
	"time"
)

/*

mCycle.New(mCycle.Opt{
	Func:      LogInit,
	SleepTime: time.Hour * 24,  // 每24 小时 执行一次
}).Start()

*/

type Cycle struct {
	Func      func()
	SleepTime time.Duration
	Status    int
}

type Opt struct {
	Func      func()
	SleepTime time.Duration
}

func New(param Opt) *Cycle {
	var CycleObj Cycle

	CycleObj.Func = param.Func

	if CycleObj.Func == nil {
		CycleObj.Func = func() {}
	}

	CycleObj.SleepTime = param.SleepTime
	CycleObj.Status = 1 // 表示开始循环

	return &CycleObj
}

func (Cy *Cycle) End() *Cycle {
	Cy.Status = 2

	return Cy
}

func (Cy *Cycle) Start() *Cycle {
	Cy.Func()
	go func() {
		for {
			if Cy.Status == 2 {
				break
			}

			time.Sleep(Cy.SleepTime) // 间隔多久执行一次
			Cy.Func()
		}
	}()

	return Cy
}
