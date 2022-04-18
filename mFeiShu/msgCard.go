package mFeiShu

import (
	"time"

	"github.com/EasyGolang/goTools/mStr"
	jsoniter "github.com/json-iterator/go"
)

type Msg_1_cardType struct {
	Config   Msg_1_Config     `json:"config"`
	Elements []Msg_1_Elements `json:"elements"`
}

type Msg_1_Config struct {
	WideScreenMode bool `json:"wide_screen_mode"`
}

type Msg_1_Elements struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

func Msg_1_card(MsgContent string) string {
	time := time.Now().Format("2006-01-02 15:04:05")

	/*
		"交易方向: **", Dir, "** \n",
		"交易币种: **", insID, "** \n",
	*/

	str := mStr.Join(
		MsgContent, "\n",
		"Time: ", time, " \n",
	)

	Config := Msg_1_Config{
		WideScreenMode: true,
	}
	elm := Msg_1_Elements{
		Tag:     "markdown",
		Content: str,
	}

	Elements := []Msg_1_Elements{}
	Elements = append(Elements, elm)

	card := Msg_1_cardType{
		Config:   Config,
		Elements: Elements,
	}
	cardStr, _ := jsoniter.Marshal(card)

	return string(cardStr)
}
