package main

import (
	"fmt"

	"github.com/Tsha-en/Calculator_App/internal/application"
)

func main() {
	fmt.Println(`Сервер запущен...
	Пожалуйста, откройте другой терминал и введите команду
	curl -X POST http://localhost/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"1+1\"}"`)
	if application.RunServer() != nil {
		fmt.Println("Ошибка, не удалось запустить программу :((")
	}

	application.RunServer()

}
