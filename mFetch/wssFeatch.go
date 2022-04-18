package mFetch

import (
	"fmt"
	"time"

	"github.com/EasyGolang/goTools/mStr"
	"github.com/gorilla/websocket"
)

type Wss struct {
	Url        string
	Conn       *websocket.Conn
	Ticker     *time.Ticker
	PongTicker *time.Ticker

	Module string      // 路径参数
	Chan   chan string // 关闭通知

	// 返回事件
	Event func(string, string)

	PingSec  int //  = 20 // 多少秒无响应则发送 ping
	NoResSec int // = 60 // 无响应则重启
}

// 写入内容
func (w *Wss) WriteData(content []byte) *Wss {
	err := w.Conn.WriteMessage(websocket.TextMessage, content)
	w.Event("Write", string(content))
	if err != nil {
		w.Event("WriteErr", mStr.ToStr(err))
		w.Close()
	}
	return w
}

// 发送 ping
func (w *Wss) SendPing() {
	err := w.Conn.WriteMessage(websocket.TextMessage, []byte("ping"))
	w.Event("Ping", mStr.ToStr(w.PingSec))
	if err != nil {
		w.Event("PingErr", mStr.ToStr(err))
		w.Close()
	}
}

// 读取数据
func (w *Wss) ReadData(outcome func(msg []byte)) {
	for {
		_, data, err := w.Conn.ReadMessage()
		if err != nil {
			w.Event("ReadErr", mStr.ToStr(err))
			w.Close()
			break
		}

		if len(data) > 7 {
			// 读取数据
			w.Event("Read", string(data))
			outcome(data)
		} else {
			// 数据大小不对
			w.Event("Pong", string(data))
		}

		w.Ticker.Reset(time.Second * time.Duration(w.PingSec))
		w.PongTicker.Reset(time.Second * time.Duration(w.NoResSec))

		// 读到关闭信号
		_, ok := <-w.Chan
		if !ok {
			break
		}
	}
}

func (w *Wss) Close() *Wss {
	w.Event("WssClose", w.Module)

	err := w.Conn.Close()
	if err != nil {
		w.Event("CloseErr", mStr.ToStr(err))
	}

	w.Ticker.Stop()
	w.PongTicker.Stop()

	// 通道未关闭时写入
	_, ok := <-w.Chan
	if ok {
		w.Chan <- "close"
	}

	close(w.Chan)
	return w
}

type WssOpt struct {
	Origin   string
	Path     string
	Event    func(string, string)
	Module   string
	PingSec  int
	NoResSec int
}

func NewWss(opt WssOpt) *Wss {
	var w Wss

	if len(opt.Path) < 2 {
		errStr := fmt.Errorf("缺少 Path 参数")
		panic(errStr)
	}

	if opt.Event == nil {
		w.Event = func(s1, s2 string) {}
	} else {
		w.Event = opt.Event
	}

	if w.PingSec < 1 {
		w.PingSec = 20
	} else {
		w.PingSec = opt.PingSec
	}

	if w.PingSec < 1 {
		w.NoResSec = 60
	} else {
		w.NoResSec = opt.NoResSec
	}

	w.Url = opt.Origin + opt.Path
	w.Module = opt.Module

	w.Ticker = time.NewTicker(time.Second * time.Duration(w.PingSec))
	w.PongTicker = time.NewTicker(time.Second * time.Duration(w.NoResSec))

	w.Chan = make(chan string)

	conn, _, err := websocket.DefaultDialer.Dial(w.Url, nil)
	w.Conn = conn

	if err != nil {
		w.Event("ConnectErr", mStr.ToStr(err))
		w.Close()
	}

	go func() {
		for range w.Ticker.C {
			w.SendPing()

			// 读到关闭信号
			_, ok := <-w.Chan
			if !ok {
				break
			}
		}
	}()

	// 如果未响应，则关闭连接重启
	go func() {
		for range w.PongTicker.C {
			w.Event("PongErr", fmt.Sprint(w.NoResSec)+"秒无响应")
			w.Close()
			break
		}
	}()

	return &w
}
