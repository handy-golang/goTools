package mWss

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type Wss struct {
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
	w.Log("Write", string(content))
	if err != nil {
		w.Log("WriteErr", err)
		w.Close()
	}
	return w
}

// 发送 ping
func (w *Wss) SendPing() {
	err := w.Conn.WriteMessage(websocket.TextMessage, []byte("ping"))
	w.Log("Ping", w.PingSec)
	if err != nil {
		w.Log("PingErr", err)
		w.Close()
	}
}

// 读取数据
func (w *Wss) ReadData(outcome func(msg []byte)) {
	for {
		_, data, err := w.Conn.ReadMessage()
		if err != nil {
			w.Log("ReadErr", err)
			w.Close()
			break
		}

		if len(data) > 7 {
			// 读取数据
			w.Event("Read", string(data))
			outcome(data)
		} else {
			// 数据大小不对
			w.Log("Pong", string(data))
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
	w.Log("WssClose", w.Module)

	err := w.Conn.Close()
	if err != nil {
		w.Log("CloseErr", err)
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

type WssParam struct {
	Url      string
	Event    func(string, string)
	Module   string
	PingSec  int
	NoResSec int
}

func Connection(param WssParam) *Wss {
	var w Wss

	w.Event = param.Event
	w.Module = param.Module

	w.PingSec = param.PingSec
	if w.PingSec < 1 {
		w.PingSec = 20
	}

	w.NoResSec = param.NoResSec
	if w.PingSec < 1 {
		w.NoResSec = 60
	}

	w.Ticker = time.NewTicker(time.Second * time.Duration(w.PingSec))
	w.PongTicker = time.NewTicker(time.Second * time.Duration(w.NoResSec))

	w.Chan = make(chan string)

	conn, _, err := websocket.DefaultDialer.Dial(param.Url, nil)
	w.Conn = conn

	if err != nil {
		w.Log("ConnectErr", err)
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
			w.Log("PongErr", fmt.Sprint(w.NoResSec)+"秒无响应", err)
			w.Close()
			break
		}
	}()

	return &w
}
