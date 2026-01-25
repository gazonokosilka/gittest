package main

import (
	"fmt"
	httpserver "gitprc/http_server"
)

func main() {
	fmt.Println("Запуск сервера")
	fmt.Println("Test")

	err := httpserver.StartHTTPServer()
	if err != nil {
		fmt.Println("Во время работы сервера произошла ошибка")
	} else {
		fmt.Println("Все гуд")
	}
}
