package api

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
	Value string `json:"value"`
}

// Encrypt encrypts the value from the Data object.
// Reference:
// https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes/.
// https://golang.org/src/crypto/cipher/example_test.go
func (d *Data) Encrypt(key string) (string, error) {
	keyBytes := []byte(key)
	valueBytes := []byte(d.Value)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// random number generated
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, valueBytes, nil)

	return fmt.Sprintf("%x", ciphertext), nil
}

// Decrypt decrypts the value from the Data object.
// Reference:
// https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes/.
// https://golang.org/src/crypto/cipher/example_test.go
func (d *Data) Decrypt(key string) (string, error) {
	keyBytes := []byte(key)
	valueBytes, err := hex.DecodeString(d.Value)
	if err != nil {
		return "", fmt.Errorf("value does not look to be encrypted: %s", err)
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()

	// Get random number and value
	nonce, ciphertext := valueBytes[:nonceSize], valueBytes[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", plaintext), nil
}
