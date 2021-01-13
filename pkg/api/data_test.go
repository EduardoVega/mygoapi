package api

import (
	"strings"
	"testing"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		Key           string
		Value         string
		ExpectedError string
		ExpectError   bool
	}{
		{
			"badSizekey",
			"mysecret",
			"invalid key size",
			true,
		},
		{
			"Thisismyverysecurekey1234567890$",
			"mysecret",
			"",
			false,
		},
	}

	for _, test := range tests {
		d := Data{test.Value}
		_, err := d.Encrypt(test.Key)

		if test.ExpectError {
			if !strings.Contains(err.Error(), test.ExpectedError) {
				t.Errorf("Returned error was incorrect, got: %v want: %v", err.Error(), test.ExpectedError)
			}
		} else {
			if err != nil {
				t.Errorf("Returned result was incorrect, got error: %v", err.Error())
			}
		}
	}
}

func TestDecrypt(t *testing.T) {
	tests := []struct {
		Key                    string
		Value                  string
		ExpectedDecryptedValue string
		ExpectedError          string
		ExpectError            bool
	}{
		{
			"badSizekey",
			"a35c18b47c54630fe35ec490fb3a1ff5db9934bbcb02ff4ec18eb96a1a1fafef3016d9e1",
			"",
			"invalid key size",
			true,
		},
		{
			"Thisismyverysecurekey1234567890$",
			"a35c18b47c54630fe35ec490fb3a1ff5db9934bbcb02ff4ec18eb96a1a1fafef3016d9e1",
			"mysecret",
			"",
			false,
		},
	}

	for _, test := range tests {
		d := Data{test.Value}
		v, err := d.Decrypt(test.Key)

		if test.ExpectError {
			if !strings.Contains(err.Error(), test.ExpectedError) {
				t.Errorf("Returned error was incorrect, got: %v want: %v", err.Error(), test.ExpectedError)
			}
		} else {
			if err != nil {
				t.Errorf("Returned result was incorrect, got error: %v", err.Error())
			}

			if v != test.ExpectedDecryptedValue {
				t.Errorf("Returned result was incorrect, got: %v want: %v", v, test.ExpectedDecryptedValue)
			}
		}
	}
}
