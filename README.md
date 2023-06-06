# goTools

一个 go 语言工具库

包含很多的工具跟模块,让代码编写更加容易

```bash
go get -u github.com/EasyGolang/goTools
```

## 包的主页

https://pkg.go.dev/github.com/EasyGolang/goTools


### Decimal替代方案 cDec

#### 测试用例
```
go test ./cDec
```

#### 使用示例
```
// 将一个浮点字符串，转换成value，并进行运算求值（更多示例可以看单元测试）
floatStr1 := "0.00012345"
floatStr2 := "1.5"
a1 := cDec.NewFromString(floatStr1)
a2 := cDec.NewFromString(floatStr2)

a3 := a1.Sub(a2)
a4 := a1.Mul(a2)
a5 := a1.Add(a2)
a6 := a1.Div(a2)
...
```