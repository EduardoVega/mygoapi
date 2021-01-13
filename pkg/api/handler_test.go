package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

type ResponseError struct {
	Error string `json:"error"`
}

func TestEncryptHandler(t *testing.T) {
	tests := []struct {
		Key           string
		JSON          string
		ExpectedCode  int
		ExpectError   bool
		ExpectedError string
	}{
		{
			"12345678901234567890123456789012",
			`{"value":"myvalue"}`,
			200,
			false,
			"",
		},
		{
			"12345678901234567890123456789012",
			`{"valueNotPresent":"myvalue"}`,
			400,
			true,
			"value is required and can not be empty",
		},
		{
			"12345678901234567890123456789012",
			`{"valueEmpty":"  "}`,
			400,
			true,
			"value is required and can not be empty",
		},
		{
			"12345678901234567890123456789012",
			`{"value":1}`,
			400,
			true,
			"cannot unmarshal number into Go struct field Data.value of type string",
		},
		{
			"1234567890",
			`{"value":"myvalue"}`,
			500,
			true,
			"invalid key size",
		},
	}

	for _, test := range tests {
		h := Handler{Key: test.Key}

		var jsonStr = []byte(test.JSON)

		req, err := http.NewRequest("POST", "/api/encrypt", bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/api/encrypt", h.EncryptHandler).Methods("POST").Headers("Content-Type", "application/json")
		router.ServeHTTP(resp, req)

		if resp.Code != test.ExpectedCode {
			t.Errorf("Returned result was incorrect, got: %v want: %v", resp.Code, test.ExpectedCode)
		}

		if test.ExpectError {
			var respError ResponseError

			decoder := json.NewDecoder(resp.Body)
			if err := decoder.Decode(&respError); err != nil {
				t.Fatal(err)
			}

			if !strings.Contains(respError.Error, test.ExpectedError) {
				t.Errorf("Returned error was incorrect, got: %v want: %v", respError.Error, test.ExpectedError)
			}
		}
	}
}

func TestDecryptHandler(t *testing.T) {
	tests := []struct {
		Key           string
		JSON          string
		ExpectedCode  int
		ExpectError   bool
		ExpectedError string
	}{
		{
			"12345678901234567890123456789012",
			`{"value":"3a33d8a6a708409aa65d380f4bb3923b5aec6bf26421b73db3a41a420aa8bf3dcd2f4b"}`,
			200,
			false,
			"",
		},
		{
			"12345678901234567890123456789012",
			`{"value":"myvalue"}`,
			500,
			true,
			"value does not look to be encrypted",
		},
		{
			"12345678901234567890123456789012",
			`{"valueNotPresent":"myvalue"}`,
			400,
			true,
			"value is required and can not be empty",
		},
		{
			"12345678901234567890123456789012",
			`{"valueEmpty":"  "}`,
			400,
			true,
			"value is required and can not be empty",
		},
		{
			"12345678901234567890123456789012",
			`{"value":1}`,
			400,
			true,
			"cannot unmarshal number into Go struct field Data.value of type string",
		},
		{
			"1234567890",
			`{"value":"a35c18b47c54630fe35ec490fb3a1ff5db9934bbcb02ff4ec18eb96a1a1fafef3016d9e1"}`,
			500,
			true,
			"invalid key size",
		},
	}

	for _, test := range tests {
		h := Handler{Key: test.Key}

		var jsonStr = []byte(test.JSON)

		req, err := http.NewRequest("POST", "/api/decrypt", bytes.NewBuffer(jsonStr))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/api/decrypt", h.DecryptHandler).Methods("POST").Headers("Content-Type", "application/json")
		router.ServeHTTP(resp, req)

		if resp.Code != test.ExpectedCode {
			t.Errorf("Returned result was incorrect, got: %v want: %v", resp.Code, test.ExpectedCode)
		}

		if test.ExpectError {
			var respError ResponseError

			decoder := json.NewDecoder(resp.Body)
			if err := decoder.Decode(&respError); err != nil {
				t.Fatal(err)
			}

			if !strings.Contains(respError.Error, test.ExpectedError) {
				t.Errorf("Returned error was incorrect, got: %v want: %v", respError.Error, test.ExpectedError)
			}
		}
	}
}
