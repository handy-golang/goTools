package mFetch

import (
	"fmt"
	"strings"
)

/*
PongErr

ConnectErr

CloseErr

WssClose

Write

WriteErr

PingErr

ReadErr

Pong

*/
func (w *Wss) Log(lType string, s ...any) *Wss {
	str := fmt.Sprintf("模块 %+v : %+v \n", w.Module, s)
	find := strings.Contains(str, "apiKey")
	if find {
		str = ("包含登录信息，加密输出:****************")
	}

	w.Event(lType, str)

	return w
}
