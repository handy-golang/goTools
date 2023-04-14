package mEncrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/EasyGolang/goTools/mStr"
)

/*

	Golang Ase 解密

*/

// 必须要结合 https://github.com/AItrade-mo7/jsEncrypt 项目来使用

func AseDecrypt(cipherText, key string) string {
	bs, err := hex.DecodeString((cipherText))
	if err != nil {
		return ""
	}

	pKey := PaddingLeft16(key)

	block, err := aes.NewCipher(pKey)
	if err != nil {
		return ""
	}
	blockModel := cipher.NewCBCDecrypter(block, pKey)
	plantText := make([]byte, len(bs))
	blockModel.CryptBlocks(plantText, bs)
	plantText = PKCS7UnPadding(plantText)
	return mStr.ToStr(plantText)
}

func PKCS7UnPadding(plantText []byte) []byte {
	length := len(plantText)
	unpadding := int(plantText[length-1])
	if length-unpadding < 0 {
		return []byte("")
	}
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
