package mGin

/*
var (
	// OK
	OK      = response(200, "Succeed")           // 通用成功
	OKNoMsg = response(200, "")                  // 无提示的成功
	Err     = response(500, "ServerErr")         // 通用错误
	Limit   = response(201, "too many requests") // 请求频率过快

	// 模块级错误码 - 用户模块
	LoginSucceed    = response(200, "登录成功")
	ErrUserName     = response(202, "用户名不正确")
	ErrUserPassword = response(202, "用户名或密码不正确")

	//  订单模块
	OrderSucceed = response(200, "下单成功")
	OrderErr     = response(301, "下单失败")

	// K线模块
	Kerr = response(210, "K线错误")

	// SwitchCoin
	SwitchSucceed = response(200, "切换币种成功")

	// SetConfig
	ConfigSucceed = response(200, "设置成功")
	ConfigErr     = response(205, "持仓时禁止切换")
)

mGin.ConfigSucceed.WithData(result)

*/
type ResType struct {
	Code int    `json:"Code"` // 返回码
	Data any    `json:"Data"` // 返回数据
	Msg  string `json:"Msg"`  // 描述
}

func Response(code int, msg string) *ResType {
	return &ResType{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func (o *ResType) WithMsg(message string) ResType {
	return ResType{
		Code: o.Code,
		Msg:  message,
		Data: o.Data,
	}
}

func (o *ResType) WithData(data any) ResType {
	return ResType{
		Code: o.Code,
		Msg:  o.Msg,
		Data: data,
	}
}
