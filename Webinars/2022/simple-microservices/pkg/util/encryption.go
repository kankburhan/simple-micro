package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func AESEncryption(text string, key string) (string, error) {

	textByte := []byte(text)
	keyByte := []byte(key)

	c, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	encoded := base64.StdEncoding.EncodeToString(gcm.Seal(nonce, nonce, textByte, nil))

	return encoded, err
}

func AESDecryption(cipherText string, key string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	keyByte := []byte(key)
	cipherTextByte := []byte(decoded)
	c, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return "", err
	}

	nonce, cipherTextByte := cipherTextByte[:nonceSize], cipherTextByte[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherTextByte, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), err
}
