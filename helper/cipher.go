package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// 適宜変更してください
var iv = []byte{71, 84, 54, 46, 114, 113, 107, 50, 116, 106, 50, 45, 54, 101, 69, 66}

// 適宜変更してください
var encryptkey = []byte{56, 112, 100, 52, 70, 117, 90, 118, 87, 117, 72, 115, 71, 121, 76, 100}

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// 暗号化
func TextEncrypt(plaintext []byte) (string, error) {
	c, err := aes.NewCipher(encryptkey)
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBEncrypter(c, iv)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return encodeBase64(ciphertext), nil
}

// 復号化
func TextDecrypt(encodeciphertext string) (string, error) {
	c, err := aes.NewCipher(encryptkey)
	if err != nil {
		return "", err
	}
	ciphertext := decodeBase64(encodeciphertext)
	cfbdec := cipher.NewCFBDecrypter(c, iv)
	plaintext := make([]byte, len(ciphertext))
	cfbdec.XORKeyStream(plaintext, ciphertext)
	return string(plaintext), nil
}
