package routes

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"net/http"
	"strconv"
)

type JsonModel struct {
	Key string `json:"key"`
	Val int    `json:"val"`
}

func JsonSum(database *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var jm JsonModel
			result := make(map[string]int)

			decodeErr := json.NewDecoder(r.Body).Decode(&jm)
			if decodeErr != nil {
				http.Error(w, "Body decoding problem", http.StatusInternalServerError)
				return
			}

			// Если в базе есть запись под данным ключем, то
			if database.Exists(jm.Key).Val() != 0 {
				stringValueFromDb, getErr := database.Get(jm.Key).Result()
				if getErr != nil {
					fmt.Fprintln(w, "Error while getting data from database", getErr)
					return
				}
				//конвертируем строчное значение из базы в числовое
				intValueFromDb, convertErr := strconv.Atoi(stringValueFromDb)
				if convertErr != nil {
					fmt.Fprintln(w, "Error during conversion", convertErr)
					return
				}
				//складываем значение из базы со значением отправленном в json
				intValueFromDb = jm.Val + intValueFromDb

				//обновляем запись
				setErr := database.Set(jm.Key, intValueFromDb, 0).Err()
				if setErr != nil {
					fmt.Fprintln(w, "Error while setting data to database", setErr)
					return
				}
				//приводим получившиеся значения под мапу, для вывода значения в соответствии с заданием
				result[jm.Key] = intValueFromDb
				returnValue, marshalErr := json.Marshal(result)
				if marshalErr != nil {
					fmt.Fprintln(w, "Error during serialization", marshalErr)
					return
				}

				fmt.Fprintf(w, string(returnValue))
			} else {
				//в случае, если записи под отправленным ключем не находится, записываем в бд
				setErr := database.Set(jm.Key, jm.Val, 0).Err()

				if setErr != nil {
					fmt.Fprintln(w, "Error while setting data to database", setErr)
					return
				}
				result[jm.Key] = jm.Val

				returnValue, marshalErr := json.Marshal(result)
				if marshalErr != nil {
					fmt.Fprintln(w, "Error during serialization", marshalErr)
					return
				}

				fmt.Fprintf(w, "Value %s was added to database", returnValue)
			}
		default:
			http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}
