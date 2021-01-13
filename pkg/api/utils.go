package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

// EmptyString validates if a string is empty.
// Could be improved by using something like https://github.com/xeipuuv/gojsonschema.
func EmptyString(value string) bool {
	r, _ := regexp.Compile("^\\s*$")

	if r.MatchString(value) {
		return true
	}

	return false
}

// CreateRandomKey creates a random key for encryp and decrypt operations.
// This will be the default if no key is found in /etc/mygoapi.
func CreateRandomKey() string {
	log.Println("creating default random key")

	uuid := uuid.New().String()

	return uuid[:len(uuid)-4]
}

// ReadKeyFromFile reads the key for encryption and decryption from a file.
func ReadKeyFromFile(r io.Reader) (string, error) {
	log.Println("reading key from file")

	content, err := ioutil.ReadAll(r)
	if err != nil {
		// any error here will not affect the app.
		// a default key will be generated.
		return "", err
	}

	strContent := strings.TrimSuffix(string(content), "\n")

	if len(strContent) != 32 {
		return "", fmt.Errorf("error key length must be of 32 characters. Current size is: %d", len(strContent))
	}

	return string(strContent), nil
}
