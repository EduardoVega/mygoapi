package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"math/rand"
	"mygoapi/pkg/handler"
	"net/http"

	"github.com/gorilla/mux"
)

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Key holds the passphrase used for encryption and decryption.
var Key string

func main() {

	h := handler.Handler{Key: Key}
	router := mux.NewRouter()

	router.HandleFunc("/api/encrypt", h.EncryptHandler).Methods("POST").Headers("Content-Type", "application/json")
	router.HandleFunc("/api/decrypt", h.DecryptHandler).Methods("POST").Headers("Content-Type", "application/json")

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8081", router))
}

func init() {
	Key = CreateKeyHash(CreateRandomKey())
}

// CreateRandomKey creates a random key for encryp and decrypt operations.
// This will be the default if no key is found as file or env var.
// Not the best solution but it works.
func CreateRandomKey() string {
	b := make([]byte, 32)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}
	return string(b)
}

// CreateKeyHash creates a hash from the key to meet length criteria for AES (32 character key).
// Used mainly when the user provides the key to avoid any problem.
func CreateKeyHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// TODO
// Read from file
// Read from env var
