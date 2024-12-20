package main

import (
	"fmt"

	"github.com/Tsha-en/Calculator_App/internal/application"
)

func main() {
	app := application.New()
	//app.Run()
	if app.RunServer() != nil {
		fmt.Println("Ошибка, не удалось запустить программу.((")
	}
	app.RunServer()
}
