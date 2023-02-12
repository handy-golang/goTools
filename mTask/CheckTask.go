package mTask

import (
	"fmt"
	"strings"

	"github.com/EasyGolang/goTools/mStr"
)

// 注册模板
var TaskTypeList = []string{
	"SendEmail",
}

type TaskType struct {
	TaskID        string         `bson:"TaskID"`         // 任务ID  当做文件名字
	TaskType      string         `bson:"TaskType"`       // 任务类型 & 解析模板
	Content       map[string]any `bson:"Content"`        // 任务内容 需要使用不同的模板去解析
	Source        string         `bson:"Source"`         // 任务来源
	Description   string         `bson:"Description"`    // 任务描述
	CreateTime    int64          `bson:"CreateTime"`     // 任务创建时间 由 任务发起方生成
	CreateTimeStr string         `bson:"CreateTimeStr"`  // 任务创建时间
	EndTime       int64          `bson:"EndTime"`        // 任务结束时间 由 任务处理方生成
	EndTimeStr    string         `bson:"EndTimeUnixStr"` // 任务结束时间
}

func CheckTask(opt TaskType) (resData TaskType, resErr error) {
	resData = TaskType{}
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

	if !find || TaskTypeNow == sepChart {
		resErr = fmt.Errorf("opt.TaskType不存在")
		return
	}

	resData = opt

	return
}
