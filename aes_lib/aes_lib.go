package aes_lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

func GenerateKey() []byte {

	randomBytes := make([]byte, 32)
	numberBytesRead, err := rand.Read(randomBytes)

	if err != nil {
		log.Fatal("Error generating key")
	}

	if numberBytesRead != 32 {
		log.Fatal("Error generating 32 bytes key => ", err)
	}

	return randomBytes
}

func EncryptData(key []byte, data []byte) []byte {

	block, err := aes.NewCipher(key)

	if err != nil {
		log.Fatal("Error creating block => ", err)
	}

	cipherText := make([]byte, aes.BlockSize+len(data))
	iv := cipherText[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)

	if err != nil {
		log.Fatal(err)
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], data)

	return cipherText

}

func DecryptData(key, cipherText []byte) []byte {

	block, err := aes.NewCipher(key)

	if err != nil {
		log.Fatal(err)
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(cipherText, cipherText)

	return cipherText
}
