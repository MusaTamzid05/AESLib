package operation

import (
	"io/ioutil"
	"log"
	"security_lib/aes_lib/aes_lib"
)

func GenerateKey(path string) {
	log.Println("[*] Generating key")
	key := aes_lib.GenerateKey()
	log.Println("[*] Saving key to ", path)
	err := ioutil.WriteFile(path, key, 0644)

	if err != nil {
		log.Println("[=] Error saving the key file to ", path)
		log.Println(err)
	}

}

func Encrypt(keyFilePath, filePath, dstPath string) {
	log.Println("[*] Encrypting ", filePath)
	keyData, err := ioutil.ReadFile(keyFilePath)

	if err != nil {
		log.Println("[=] Error reading ", keyFilePath)
		log.Fatalln(err)
	}

	fileData, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Println("[=] Error reading ", keyFilePath)
		log.Fatalln(err)

	}

	enctyptedData := aes_lib.EncryptData(keyData, fileData)

	err = ioutil.WriteFile(dstPath, enctyptedData, 0644)

	if err != nil {
		log.Println("[=] Error saving => ", dstPath)
		log.Println(err)
	}

	log.Println("[*] Encrypting save to ", dstPath)
}

func Decrypt(keyFilePath, filePath, dstPath string) {
	log.Println("[*] Decrypt ", filePath)
	keyData, err := ioutil.ReadFile(keyFilePath)

	if err != nil {
		log.Println("[=] Error reading ", keyFilePath)
		log.Fatalln(err)
	}

	fileData, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Println("[=] Error reading ", keyFilePath)
		log.Fatalln(err)

	}

	decryptedData := aes_lib.DecryptData(keyData, fileData)

	log.Println("[*] Writting decrypt data to ", dstPath)
	err = ioutil.WriteFile(dstPath, decryptedData, 0644)

	if err != nil {
		log.Println("[=] Error saving => ", dstPath)
		log.Println(err)
	}

}
