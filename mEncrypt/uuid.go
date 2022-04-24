package mEncrypt

import (
	"strings"

	"github.com/google/uuid"
)

func GetUUID() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}
