package mOKX

type ReqType struct {
	Code string `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}
