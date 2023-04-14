package mEncrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/EasyGolang/goTools/mStr"
)

// 必须要结合 https://github.com/AItrade-mo7/jsEncrypt 项目来使用
// Decrypt Golang解密
// ciphertext  important important 上面js的生成的密文进行了 hex.encoding 在这之前必须要进行 hex.Decoding
// 上面js代码最后返回的是16进制
// 所以收到的数据hexText还需要用hex.DecodeString(hexText)转一下,这里略了
func Decrypt(cipherText, key string) string {
	bs, err := hex.DecodeString((cipherText))
	if err != nil {
		return ""
	}

	pKey := PaddingLeft16(key) // 和js的key补码方法一致

	block, err := aes.NewCipher(pKey) // 选择加密算法
	if err != nil {
		return ""
	}
	blockModel := cipher.NewCBCDecrypter(block, pKey) // 和前端代码对应:   mode: CryptoJS.mode.CBC,// CBC算法
	plantText := make([]byte, len(bs))
	blockModel.CryptBlocks(plantText, bs)
	plantText = PKCS7UnPadding(plantText) // 和前端代码对应:  padding: CryptoJS.pad.Pkcs7
	return mStr.ToStr(plantText)
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	return plantText[:(length - unpadding)]
}

func PaddingLeft16(str string) []byte {
	pKey := mStr.ToStr(str)
	if len(pKey) < 16 {
		tL := len(pKey)
		for i := 0; i < 16-tL; i++ {
			pKey += "0"
		}
	} else {
		pKey = pKey[16:]
	}
	return []byte(pKey)
}
