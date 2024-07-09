// source: https://gist.github.com/fracasula/38aa1a4e7481f9cedfa78a0cdd5f1865

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

var (
	// 32 bytes long secret
	secretKey string = "N1PCdw3M2B1TfJhoaY2mL736p2vCUc47"
)

func Encrypt(message string) (string, error) {
	byteMsg := []byte(message)

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}

	// buffer the same length as plaintext
	cipherText := make([]byte, aes.BlockSize+len(byteMsg))

	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("could not encrypt: %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(byteMsg))

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decrypt(message string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", fmt.Errorf("could not base64 decode: %v", err)
	}

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("could not create new cipher: %v", err)
	}

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("invalid cipher text block size")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}
