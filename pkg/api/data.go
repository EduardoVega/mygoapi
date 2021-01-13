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
// Reference: https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes/.
// I did nothing (copy and paste). At least I can explain what this is doing.
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

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, valueBytes, nil)

	return fmt.Sprintf("%x", ciphertext), nil
}

// Decrypt decrypts the value from the Data object.
// Reference: https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes/.
// I did nothing (copy and paste). At least I can explain what this is doing.
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

	nonce, ciphertext := valueBytes[:nonceSize], valueBytes[nonceSize:]

	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", plaintext), nil
}

// TODO
// Understant what the code does
