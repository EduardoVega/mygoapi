package object

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

// Data holds the JSON document received from the encrypt and decrypt apis.
type Data struct {
	Value string
}

// Encrypt encrypts the value from the Data object.
// Reference: https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes/
func (d *Data) Encrypt(key string) (string, error) {
	keyBytes, _ := hex.DecodeString(key)
	valueBytes := []byte(d.Value)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, valueBytes, nil)

	return fmt.Sprintf("%x", ciphertext), nil
	// return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the value from the Data object.
// Reference: https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes/
func (d *Data) Decrypt(key string) (string, error) {
	keyBytes, _ := hex.DecodeString(key)
	valueBytes, _ := hex.DecodeString(d.Value)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := valueBytes[:nonceSize], valueBytes[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext), nil
}
