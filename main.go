package main

import (
	"fmt"

	"github.com/EasyGolang/goTools/mFeiShu"
	"github.com/EasyGolang/goTools/mStr"
)

func main() {
	fmt.Println(" =========  START  ========= ")

	feishuApp := mFeiShu.New(mFeiShu.Opt{
		AppID:     "cli_a28394cb5478d00d",
		AppSecret: "MDMJs33KsiH9FAxr74MqSXG3lTL4ptPT",
	})

	str := mStr.Join(
		"交易方向: **", "开多", "** \n",
		"交易币种: **", "avax", "** \n",
	)

	feishuApp.SendMessage(mFeiShu.MsgOpt{
		ReceiveType: "user_id",
		ReceiveId:   "d81242gc",
		Content:     str,
	})

	fmt.Println(" =========   END   ========= ")
}
