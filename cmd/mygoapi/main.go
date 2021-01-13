package main

import (
	"log"
	"mygoapi/pkg/api"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// characters is used to generate a random key.
const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// keyFile holds the path to the file containing the key.
const keyFile = "/etc/mygoapi"

// key holds the passphrase used for encryption and decryption.
var key string

func main() {

	h := api.Handler{Key: key}
	router := mux.NewRouter()

	router.HandleFunc("/api/encrypt", h.EncryptHandler).Methods("POST").Headers("Content-Type", "application/json")
	router.HandleFunc("/api/decrypt", h.DecryptHandler).Methods("POST").Headers("Content-Type", "application/json")

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8081", router))
}

func init() {
	// Generate default key.
	key = api.CreateRandomKey()

	// Get key from file.
	f, err := os.Open(keyFile)
	defer f.Close()
	if err != nil {
		log.Printf("warning opening file with key: %q. Random key will be used", err)
		return
	}

	keyFromFile, err := api.ReadKeyFromFile(f)
	if err != nil {
		log.Printf("warning reading key from file: %q. Random key will be used", err)
		return
	}

	key = keyFromFile
}
