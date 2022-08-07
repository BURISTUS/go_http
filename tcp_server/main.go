package main

import (
	"bufio"
	"net"
	"strconv"
	"strings"
)
import "fmt"

const (
	firstElOfArr  = 0
	secondElOfArr = 1
)

func main() {
	var result []int
	var stringSlice []string
	var updatedString strings.Builder
	buf := make([]byte, 1024)

	fmt.Println("Launching server...")
	//Установка прослушивателя порта
	ln, lnErr := net.Listen("tcp", ":8081")
	if lnErr != nil {
		fmt.Println("Problem during listening", lnErr)
		return
	}
	//Открытие порта
	conn, connErr := ln.Accept()
	if connErr != nil {
		fmt.Println("Problem during port opening", connErr)
		return
	}
	//Прослушивание сообщения
	bufio.NewReader(conn).Read(buf)
	stringSlice = strings.Split(string(buf), "\r\n")
	//Приведение полученной строки к нужному виду
	for i := 0; i < len(stringSlice)-2; i++ {
		res := strings.Split(stringSlice[i], ",")
		var sliceOfMultVals []int

		for i := 0; i < len(res); i++ {
			val, _ := strconv.Atoi(res[i])
			sliceOfMultVals = append(sliceOfMultVals, val)
		}
		//Заполнение среза с результатами перемножений
		result = append(result, sliceOfMultVals[firstElOfArr]*sliceOfMultVals[secondElOfArr])
	}
	//Приведение строки к нужному виду
	for i := 0; i < len(result); i++ {
		value := strconv.Itoa(result[i])
		updatedString.WriteString(value + "\r\n")
	}
	//Отправление обновленной строки на строну клиента
	conn.Write([]byte(updatedString.String() + "\r\n"))
}
