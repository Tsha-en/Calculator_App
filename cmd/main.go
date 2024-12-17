package main

import (
	"Calculator_App/pkg/calculator"
	"fmt"
)

func main() {
	var expr string
	fmt.Print("Введите выражение: ")
	fmt.Scan(&expr)

	result, err := calculator.Calc(expr)
	if err != nil {
		fmt.Printf("Error calculating '%s': %s\n", expr, err)
	} else {
		fmt.Printf("Result of '%s': %f\n", expr, result)
	}
}
