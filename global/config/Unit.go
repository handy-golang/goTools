package config

import "github.com/EasyGolang/goTools/mStr"

// 计价的锚定货币
var Unit = "USDT"

var SPOT_suffix = mStr.Join("-", Unit)

var SWAP_suffix = mStr.Join(SPOT_suffix, "-SWAP")
