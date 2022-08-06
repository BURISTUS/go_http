package routes

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
)

type SecretData struct {
	key string `json:"s"`
	val int    `json:"key"`
}

func Test2(w http.ResponseWriter, r *http.Request) {
	var d map[string]interface{}

	decoderErr := json.NewDecoder(r.Body).Decode(&d)
	if decoderErr != nil {
		fmt.Println("decoderErr", decoderErr)
	}

	keys := []string{}

	for key, _ := range d {
		keys = append(keys, key)
	}

	h := hmac.New(sha512.New, []byte(keys[0]))

	// Write Data to it
	//h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	fmt.Println("Result: " + sha)
}
