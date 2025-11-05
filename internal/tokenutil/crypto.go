package tokenutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Hash data string single line with method sha256
func hashSha256(data string) (result []byte) {
	keysHash := sha256.New()
	keysHash.Write([]byte(data))
	result = keysHash.Sum(nil)

	return
}

// decode data with method base64
func decode(s string) []byte {
	data, _ := base64.StdEncoding.DecodeString(s)
	return data
}

// Remove the padding data
func unpad(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

// Adding padding to data
func pad(data []byte) []byte {
	padding := aes.BlockSize - (len(data) % aes.BlockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// Encryption data with AES-256-CBC method
func Encrypt(secretKey string, data []byte) (string, error) {
	data = pad(data)
	privateKey := hashSha256(secretKey)
	block, err := aes.NewCipher(privateKey)
	if err != nil {
		return "", err
	}

	if len(data)%aes.BlockSize != 0 {
		return "", errors.New("data must be multiple of block size")
	}

	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := privateKey[:aes.BlockSize]
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], data)

	encryptedData := append(iv, ciphertext[aes.BlockSize:]...)
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

// Decode data with AES-256-CBC Method
func Decrypt(secretKey, data string) (result string, err error) {
	key := hashSha256(secretKey) // 32 bytes for AES-256
	encryptedDataByte := decode(data)
	encryptedDataByte = encryptedDataByte[aes.BlockSize:]
	iv := key[:aes.BlockSize]
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(encryptedDataByte)%aes.BlockSize != 0 {
		return "", errors.New("invalid ciphertext length")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	decyptedDataByte := make([]byte, len(encryptedDataByte))
	mode.CryptBlocks(decyptedDataByte, encryptedDataByte)

	// Unpad the decrypted data
	return string(unpad(decyptedDataByte)), nil
}

func Format(text string, expire int64, nonce string) string {
	return fmt.Sprintf("%s:%d:%s", text, expire, nonce)
}

func ParseFormat(data string) (text string, expire int64, nonce string, err error) {
	dataSlice := strings.Split(data, ":")
	if len(dataSlice) != 3 {
		err = errors.New("invalid token format")
		return
	}

	text = dataSlice[0]
	expire, err = strconv.ParseInt(dataSlice[1], 10, 64)
	if err != nil {
		return
	}
	nonce = dataSlice[2]
	return
}
