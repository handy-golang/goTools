package mFetch

import "fmt"

func (o *Http) Log(lType string, s ...any) {
	str := fmt.Sprintf("请求地址 %+v \n %+v \n %+v \n", o.Url, s, string(o.Data))

	fmt.Println(str)
}
