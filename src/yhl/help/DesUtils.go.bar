package help

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

const DesKey string = "yhl27ml_"

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func DesEncrypt(data, keyStr string) string {
	origData, key := []byte(data), []byte(keyStr)
	block, err := des.NewCipher(key)
	if err != nil {
		Error(err)
		return ""
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	return base64.URLEncoding.EncodeToString(crypted)
}

func DesDecrypt(data, keyStr string) string {
	crypted, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		Error(err)
		return ""
	}
	key := []byte(keyStr)
	block, err := des.NewCipher(key)
	if err != nil {
		Error(err)
		return ""
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)

	return string(origData)
}
