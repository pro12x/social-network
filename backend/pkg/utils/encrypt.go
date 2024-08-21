package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"io"
	"strings"
)

func Encrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(password)), bcrypt.DefaultCost)
	return string(bytes), err
}

func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(strings.TrimSpace(hashedPassword)), []byte(strings.TrimSpace(password)))
}

func GenerateEncryptKey() (string, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return "", errors.New("error generating key: " + err.Error())
	}
	return string(key), nil
}

func EncryptData(data string, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", errors.New("error creating cipher: " + err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.New("error creating GCM: " + err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	return string(gcm.Seal(nonce, nonce, []byte(data), nil)), nil
}

func DecryptData(cipherData string, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", errors.New("error creating cipher: " + err.Error())
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.New("error creating GCM: " + err.Error())
	}

	nonceSize := gcm.NonceSize()
	if len(cipherData) < nonceSize {
		return "", errors.New("cipher too short")
	}

	nonce, cipherData := cipherData[:nonceSize], cipherData[nonceSize:]
	data, err := gcm.Open(nil, []byte(nonce), []byte(cipherData), nil)
	if err != nil {
		return "", errors.New("error decrypting content: " + err.Error())
	}
	return string(data), nil
}
