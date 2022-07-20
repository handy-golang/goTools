package mFetch

import (
	"fmt"

	"github.com/fasthttp/websocket"
)

/*

	wssConn := mFetch.NewWss(mFetch.WssOpt{
		Url: "ws://127.0.0.1:8999/api/wss",
		Event: func(s string, a any) {
			if s == "close" || s == "err" {
				fmt.Println("出错了", mStr.ToStr(a))
			}
		},
	})

	wssConn.Write([]byte("123"))

	wssConn.Read(func(msg []byte) {
		fmt.Println("read", string(msg))
	})

*/

type WssOpt struct {
	Url   string
	Event func(string, any) // s1 = succeed , err
}
type Wss struct {
	Conn  *websocket.Conn
	Event func(string, any) // s1 = red , close , err
}

func NewWss(opt WssOpt) (_this *Wss) {
	_this = &Wss{}

	// 参数检查
	errStr := []string{}
	switch {
	case len(opt.Url) < 1:
		errStr = append(errStr, "Url")
	}
	if len(errStr) > 0 {
		errStr := fmt.Errorf("缺少参数:%+v", errStr)
		panic(errStr)
	}
	// 事件处理
	_this.Event = opt.Event
	if _this.Event == nil {
		_this.Event = func(s1 string, s2 any) {}
	}

	c, _, err := websocket.DefaultDialer.Dial(opt.Url, nil)
	_this.Conn = c

	if err != nil {
		_this.Event("err", err)
		return
	}

	return
}

func (_this *Wss) Read(callback func(msg []byte)) {
	if _this.Conn == nil {
		return
	}
	for {
		_, message, err := _this.Conn.ReadMessage()
		if err != nil {
			_this.Close(err)
			return
		}
		callback(message)
		_this.Event("read", message)
	}
}

func (_this *Wss) Close(lType any) {
	if _this.Conn == nil {
		return
	}
	_this.Event("close", lType)
	_this.Conn.Close()
}

func (_this *Wss) Write(content []byte) {
	if _this.Conn == nil {
		return
	}
	err := _this.Conn.WriteMessage(websocket.TextMessage, content)
	if err != nil {
		_this.Close(err)
		return
	}
}
