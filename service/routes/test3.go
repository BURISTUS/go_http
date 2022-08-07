package routes

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
)

type DataFromJson []struct {
	A   string
	B   string
	Key string
}

func TcpClient(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var dfj DataFromJson
		var stringForTcp strings.Builder
		var stringSlice []string
		mapResult := make(map[string]int)

		buf := make([]byte, 1024)

		decoderErr := json.NewDecoder(r.Body).Decode(&dfj)
		if decoderErr != nil {
			http.Error(w, "Body decoding problem", http.StatusInternalServerError)
			return
		}
		//Пробегаемся по слайсу структур и формируем строку
		for i := range dfj {
			stringForTcp.WriteString(dfj[i].A + "," + dfj[i].B + "\r\n")
		}
		// Подключаемся к сокету
		conn, connErr := net.Dial("tcp", "127.0.0.1:8081")
		if connErr != nil {
			fmt.Fprintln(w, "Socket connection error", connErr)
			return
		}
		// Отправляем в сокет
		fmt.Fprintf(conn, stringForTcp.String()+"\r\n")
		// Прослушиваем ответ
		bufio.NewReader(conn).Read(buf)
		stringSlice = strings.Split(string(buf), "\r\n")
		//Формируем мапу
		for i := 0; i < len(stringSlice)-2; i++ {
			val, _ := strconv.Atoi(stringSlice[i])
			mapResult[dfj[i].Key] = val
		}

		byteRes, marshalErr := json.Marshal(mapResult)
		if marshalErr != nil {
			fmt.Fprintln(w, "Error during serialization", marshalErr)
			return
		}

		fmt.Fprintf(w, string(byteRes))
	default:
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

}
