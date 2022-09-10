package mFetch

import (
	"fmt"
	"time"

	"github.com/EasyGolang/goTools/mStr"
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
	Conn       *websocket.Conn
	Ticker     *time.Ticker
	PingTicker *time.Ticker
	RunIng     bool
	Event      func(string, any) // s1 = red , close , err
}

const TickerDuration time.Duration = time.Second * time.Duration(26)

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

	_this.RunIng = true

	c, _, err := websocket.DefaultDialer.Dial(opt.Url, nil)
	if err != nil || c == nil {
		_this.Close(err)
		return
	}

	_this.Conn = c
	_this.Ticker = time.NewTicker(TickerDuration)
	_this.PingTicker = time.NewTicker((TickerDuration / 4) * 3)

	// 发送 Ping
	go func() {
		for range _this.PingTicker.C {
			if !_this.RunIng {
				break
			}
			_this.Write([]byte("ping"))
		}
	}()
	// 读到 关闭信号
	go func() {
		for range _this.Ticker.C {
			if !_this.RunIng {
				break
			}
			errStr := fmt.Errorf("长时间未响应")
			_this.Close(errStr)
		}
	}()

	return
}

func (_this *Wss) Read(callback func(msg []byte)) {
	if _this.Conn == nil {
		return
	}
	for {

		if !_this.RunIng {
			break
		}

		_, message, err := _this.Conn.ReadMessage()
		if err != nil {
			_this.Close(err)
			return
		}
		_this.Event("Read", message)
		_this.Ticker.Reset(TickerDuration)
		_this.PingTicker.Reset(TickerDuration)

		// 有回应则重置
		if mStr.ToStr(message) == "pong" {
			// pong
		} else {
			callback(message)
		}
	}
}

func (_this *Wss) Close(lType any) {
	if _this.Conn == nil {
		return
	}
	_this.Event("Close", lType)
	_this.Ticker.Stop()
	_this.PingTicker.Stop()
	_this.RunIng = false
	_this.Conn.Close()
}

func (_this *Wss) Write(content []byte) {
	if _this.Conn == nil {
		return
	}

	if !_this.RunIng {
		return
	}

	_this.Event("Write", content)
	err := _this.Conn.WriteMessage(websocket.TextMessage, content)
	if err != nil {
		_this.Close(err)
		return
	}
}
