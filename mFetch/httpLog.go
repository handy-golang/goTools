package mFetch

import "fmt"

func (o *Http) Log(lType string, s ...any) *Http {
	str := fmt.Sprintf("请求地址 %+v \n %+v \n %+v \n", o.Url, s, string(o.Data))

	o.Event(lType, str)
	return o
}
