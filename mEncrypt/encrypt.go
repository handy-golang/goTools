package mEncrypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/EasyGolang/goTools/mStr"
)

func Sha256(message string, secretKey string) (string, error) {
	mac := hmac.New(sha256.New, []byte(secretKey))
	_, err := mac.Write([]byte(message))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(mac.Sum(nil)), nil
}

func Sum256(data string) string {
	org := []byte(data)
	sum := sha256.Sum256(org)
	c := fmt.Sprintf("%x", sum)
	return c
}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

type OkxSignInfo struct {
	Timestamp   string
	Method      string
	RequestPath string
	Body        string
	SecretKey   string
}

func GetOkxSign(param OkxSignInfo) string {
	signStr := mStr.Join(
		param.Timestamp,
		strings.ToUpper(param.Method),
		param.RequestPath,
		param.Body,
	)

	sign, err := Sha256(signStr, param.SecretKey)
	if err != nil {
		return ""
	}
	return sign
}
