package main

import (
	"security_lib/aes_lib/operation"
)

func main() {
	//operation.GenerateKey("secret.key")
	//operation.Encrypt("secret.key", "m3.jpg", "encrypt.data")
	operation.Decrypt("secret.key", "encrypt.data", "dectype.jpg")
}
