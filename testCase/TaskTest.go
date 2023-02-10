package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mTask"
)

func TaskTest() {
	resData, resErr := mTask.CheckTask(mTask.TaskType{
		TaskID:   mEncrypt.GetUUID(),
		TaskType: "SendsEmail",
		Content: map[string]any{
			"jsonrpc": "2.0",
			"id":      "5",
		},
	})

	fmt.Println(resData, resErr)
}
