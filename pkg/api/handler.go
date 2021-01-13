package api

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

// Handler holds the data for the handler functions.
type Handler struct {
	Key string
}

// EncryptHandler initiates the encrypt operation.
func (h *Handler) EncryptHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("handling encrypt endpoint")

	var data Data

	// Decode JSON.
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		writeResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	// Validate value.
	if emptyString(data.Value) {
		writeResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "value is required and can not be empty",
		})
		return
	}

	// Encrypt value.
	encryptedValue, err := data.Encrypt(h.Key)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	// Write response.
	writeResponse(w, http.StatusOK, map[string]interface{}{
		"value": encryptedValue,
	})
}

// DecryptHandler initiates the decrypt operation.
func (h *Handler) DecryptHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("handling decrypt endpoint")

	var data Data

	// Decode JSON.
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		writeResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	// Validate value.
	if emptyString(data.Value) {
		writeResponse(w, http.StatusBadRequest, map[string]interface{}{
			"error": "value is required and can not be empty",
		})
		return
	}

	// Decrypt value.
	value, err := data.Decrypt(h.Key)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	// Write response.
	writeResponse(w, http.StatusOK, map[string]interface{}{
		"value": value,
	})
}

// emptyString validates if a string is empty.
// Could be improved by using something like https://github.com/xeipuuv/gojsonschema.
func emptyString(value string) bool {
	r, _ := regexp.Compile("^\\s*$")

	if r.MatchString(value) {
		return true
	}

	return false
}

// writeResponse writes API responses for the handlers.
func writeResponse(w http.ResponseWriter, statusCode int, data map[string]interface{}) {
	response, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
