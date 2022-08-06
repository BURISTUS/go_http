package routes

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
)

func Test3(w http.ResponseWriter, r *http.Request) {
	var d map[string]interface{}

	decoderErr := json.NewDecoder(r.Body).Decode(&d)
	if decoderErr != nil {
		fmt.Println("decoderErr", decoderErr)
	}

	// Подключаемся к сокету
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		// Чтение входных данных от stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// Отправляем в socket
		fmt.Fprintf(conn, text+"\n")
		// Прослушиваем ответ
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
