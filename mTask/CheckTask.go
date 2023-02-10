package mTask

import (
	"fmt"
	"strings"

	"github.com/EasyGolang/goTools/mStr"
)

var TaskTypeList = []string{
	"SendEmail",
	"爱",
	"你",
}

type TaskType struct {
	TaskID   string         `bson:"TaskID"`   // 任务ID
	TaskType string         `bson:"TaskType"` // 任务类型&解析模板  SendEmail, SendOrder 等
	Content  map[string]any `bson:"Content"`  // 任务内容 用不同的模板去解析
}

func CheckTask(opt TaskType) (resData string, resErr error) {
	resData = ""
	resErr = nil

	if len(opt.TaskID) < 20 {
		resErr = fmt.Errorf("opt.TaskID 长度不足")
		return
	}

	sepChart := "&"
	TaskTypeALl := strings.Join(TaskTypeList, sepChart)
	TaskTypeALl = mStr.Join(TaskTypeALl, sepChart)

	TaskTypeNow := mStr.Join(opt.TaskType, sepChart)

	find := strings.Contains(TaskTypeALl, TaskTypeNow)

	if !find {
		resErr = fmt.Errorf("opt.TaskType不存在")
		return
	}

	resData = opt.TaskType

	return
}
