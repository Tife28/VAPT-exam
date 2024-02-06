package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var encryptionKey = []byte("4e8f1670f502a3d40717709e5f80d67c") // This is where I inserted the crypt key

func decrypt(key []byte, ciphertext string) (string, error) {
	// Decode the Base64-encoded ciphertext
	encryptedData, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// Initialize the AES block cipher with the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create a stream cipher for decryption
	stream := cipher.NewCTR(block, key[aes.BlockSize:])

	// Decrypt the ciphertext
	plaintext := make([]byte, len(encryptedData))
	stream.XORKeyStream(plaintext, encryptedData)

	return string(plaintext), nil
}

func main() {
	ciphertext := "cb15h+Mzl5pZxeNSWe3b" // I inserted the encrypted credit card number 

	decryptedNumber, err := decrypt(encryptionKey, ciphertext)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		return
	}

	fmt.Println("Decrypted Number:", decryptedNumber)
}