package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/db"
)

type JsonModel struct {
	key string `json:"key"`
	val int    `json:"val"`
}

func Test1(database *db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var jm JsonModel
		var d map[string]int

		decodeErr := json.NewDecoder(r.Body).Decode(&jm)
		if decodeErr != nil {
			fmt.Fprintf(w, "wrong schema of json", decodeErr)
		}
		fmt.Println("jm key", jm.key)

		if database.Client.Exists(jm.key).Val() != 0 {
			value, err := database.Client.Get(jm.key).Result()
			json.Unmarshal([]byte(value), &d)

			d[jm.key] = jm.val + d[jm.key]
			res, marshalErr := json.Marshal(d)

			if marshalErr != nil {
				fmt.Println("Marshal Error", marshalErr)
			}

			writeErr := database.Client.Set(jm.key, res, 0).Err()

			if writeErr != nil {
				fmt.Println("Write error", writeErr)
			}

			if err != nil {
				fmt.Println(err)
			}

			updatedValue, err := database.Client.Get(jm.key).Result()
			fmt.Fprintln(w, updatedValue)
			fmt.Println("Значение обнавлено")
		} else {
			setErr := database.Client.Set(jm.key, jm.val, 0).Err()
			fmt.Println(jm.key)

			if setErr != nil {
				fmt.Println(setErr)
			}

			val, getErr := database.Client.Get(jm.key).Result()

			if getErr != nil {
				fmt.Println("Get error", getErr)
			}
			fmt.Println("Значение добавлено")
			fmt.Fprintln(w, val)

		}

	}
}
