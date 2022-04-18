package mCycle

import (
	"time"
)

type CycleParam struct {
	Func      func()
	SleepTime time.Duration
}

type Cycle struct {
	Func      func()
	SleepTime time.Duration
	Status    int
}

func New(param CycleParam) *Cycle {
	var CycleObj Cycle

	CycleObj.Func = param.Func
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
