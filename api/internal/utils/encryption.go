package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func DecryptData(value string, key []byte) ([]byte, error) {
	decodedCipherText, _ := base64.StdEncoding.DecodeString(value)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, cipherText := decodedCipherText[:nonceSize], value[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, []byte(cipherText), nil)
	if err != nil {
		return nil, err
	}
	return plainText, err
}

func EncryptData(value []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)
	cipherText := gcm.Seal(nonce, nonce, value, nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}
