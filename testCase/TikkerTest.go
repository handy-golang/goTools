package testCase

import "github.com/EasyGolang/goTools/mTikker"

func TikkerTest() {
	mTikker.NewTikker(mTikker.TikkerOpt{
		ShellContent: `
		mEcho "我爱你22 1231  123 "
		`,
	}).RunToPm2()
}
