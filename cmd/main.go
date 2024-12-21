package main

import (
	"fmt"

	"github.com/Tsha-en/Calculator_App/internal/application"
)

func main() {
	if application.RunServer() != nil {
		fmt.Println("Ошибка, не удалось запустить программу.((")
	}
	application.RunServer()
}
