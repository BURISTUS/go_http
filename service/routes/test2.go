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
	Val string `json:"s"`
	Key string `json:"key"`
}

func GetHashFromJson(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var sc SecretData

		decoderErr := json.NewDecoder(r.Body).Decode(&sc)
		if decoderErr != nil {
			http.Error(w, "Body decoding problem", http.StatusInternalServerError)
			return
		}

		//Создаем новый HMAC с указанием типа хэша и ключа
		h := hmac.New(sha512.New, []byte(sc.Key))
		// Записываем данные
		h.Write([]byte(sc.Val))
		//Получаем результат закодированный в виде шестнадцатеричной строки
		sha := hex.EncodeToString(h.Sum(nil))

		fmt.Fprintln(w, sha)
	default:
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}
}
