package api

import (
	"bytes"
	"strings"
	"testing"
)

func TestCreateRandomKey(t *testing.T) {
	key := CreateRandomKey()

	if len(key) != 32 {
		t.Errorf("Returned key was incorrect, got: size %d want: 32 size", len(key))
	}
}

func TestReadKeyFromFile(t *testing.T) {

	tests := []struct {
		KeyFileContent string
		ExpectedKey    string
		ExpectedError  string
		ExpectError    bool
	}{
		{
			"invalidSizeKey",
			"",
			"error key length must be of 32 characters",
			true,
		},
		{
			"Thisismyverysecurekkey1234567890",
			"Thisismyverysecurekkey1234567890",
			"",
			false,
		},
	}

	for _, test := range tests {
		buf := bytes.NewBufferString(test.KeyFileContent)
		k, err := ReadKeyFromFile(buf)

		if test.ExpectError {
			if !strings.Contains(err.Error(), test.ExpectedError) {
				t.Errorf("Returned error was incorrect, got: %v want: %v", err.Error(), test.ExpectedError)
			}
		} else {
			if k != test.ExpectedKey {
				t.Errorf("Returned result was incorrect, got: %v want: %v", k, test.ExpectedKey)
			}
		}
	}
}
