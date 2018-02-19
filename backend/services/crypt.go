package services

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strings"
)

const (
	presalt  = "R00qFyRJY0Pvfqjk3Oq9K4FXFwkfDNHyZ05ufj0dYcUb2hF79woYXhESMQSj5lN9kiAi9SW33LmAvo"
	postsalt = "BfpZXF2n3TEG7pe4EFe1GFttYxzv4CJUOpFHx08NEVWV6RpSn6KofmLA1JBrTZWJfAWbLoiaSDBc55fMx4ICS4pUSr9125xkH12Y"
//	aesKey   = "bNCKcOvcKU5wLxPVhLcF5MtExncnpkhB"
)

func GetHash(data string) string {
	h256 := sha256.New()
	out := fmt.Sprintf("%s%s%s", presalt, data, postsalt)
	io.WriteString(h256, out)

	return fmt.Sprintf("%x\n", h256.Sum(nil))
}

func GetAESEncode(data string, aesKey string) (string, error) {
	key := []byte(aesKey)
	encryptMsg, err := Encrypt(key, data)
	return encryptMsg, err
}

func GetAESDecode(data string, aesKey string) (string, error) {
	key := []byte(aesKey)
	msg, err := Decrypt(key, data)
	return msg, err
}

func Encrypt(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	msg := Pad([]byte(text))
	ciphertext := make([]byte, aes.BlockSize+len(msg))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(msg))
	finalMsg := removeBase64Padding(base64.URLEncoding.EncodeToString(ciphertext))
	return finalMsg, nil
}

func Decrypt(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	decodedMsg, err := base64.URLEncoding.DecodeString(addBase64Padding(text))
	if err != nil {
		return "", err
	}

	if (len(decodedMsg) % aes.BlockSize) != 0 {
		return "", errors.New("blocksize must be multipe of decoded message length")
	}

	iv := decodedMsg[:aes.BlockSize]
	msg := decodedMsg[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(msg, msg)

	unpadMsg, err := Unpad(msg)
	if err != nil {
		return "", err
	}

	return string(unpadMsg), nil
}

func addBase64Padding(value string) string {
	m := len(value) % 4
	if m != 0 {
		value += strings.Repeat("=", 4-m)
	}

	return value
}

func removeBase64Padding(value string) string {
	return strings.Replace(value, "=", "", -1)
}

func Pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func Unpad(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])

	if unpadding > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil
}
