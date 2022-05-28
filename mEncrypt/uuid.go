package mEncrypt

import (
	"strings"
	"time"

	"github.com/EasyGolang/goTools/mCount"
	"github.com/EasyGolang/goTools/mStr"
	"github.com/google/uuid"
)

func GetUUID() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}

func TimeID() string {
	now := time.Now().Format("060102150405")
	uuid := GetUUID()
	uuidArr := strings.Split(uuid, "")
	start := mCount.GetRound(0, int64(len(uuidArr)-3))

	sArr := uuidArr[start : start+3]

	s := strings.Join(sArr, "")

	returnStr := mStr.Join(
		now, "m", s,
	)

	return returnStr
}
