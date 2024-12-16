package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func ValidateExpression(s string) error {
	if s == "" {
		return fmt.Errorf("invalid number format: empty expression")
	}

	// Регулярное выражение для проверки валидности выражения
	if matched, _ := regexp.MatchString(`^[0-9+\-*/().\s]+$`, s); !matched {
		return fmt.Errorf("invalid number format: contains invalid characters")
	}

	return nil
}

// Поиск скобок в выражении
func FindBrackets(s string) (int, int) {
	open := -1
	close := -1
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			open = i
		} else if s[i] == ')' {
			close = i
			break
		}
	}
	return open, close
}

func MulDivRes(s string) (float64, error) {
	// Удаляем пробелы из строки
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, "")
	terms := []float64{}
	operators := []string{}
	currentNum := ""

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if isOperator(ch) {
			if currentNum != "" {
				num, err := strconv.ParseFloat(currentNum, 64)
				if err != nil {
					return 0, fmt.Errorf("invalid number format")
				}
				terms = append(terms, num)
				currentNum = ""
			}
			operators = append(operators, string(ch))
		} else {
			currentNum += string(ch)
		}
	}

	// Добавляем последнее число, если оно есть
	if currentNum != "" {
		num, err := strconv.ParseFloat(currentNum, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number format")
		}
		terms = append(terms, num)
	}

	// Проверка на наличие операторов и операндов
	if len(terms) == 0 {
		return 0, fmt.Errorf("invalid number format: no numbers found")
	}
	if len(operators) == 0 {
		return terms[0], nil
	}

	// Обрабатываем операции умножения и деления
	for i := 0; i < len(operators); i++ {
		if operators[i] == "*" || operators[i] == "/" {
			if i >= len(terms)-1 {
				return 0, fmt.Errorf("invalid number format: not enough operands")
			}
			if operators[i] == "*" {
				terms[i] = terms[i] * terms[i+1]
			} else {
				if terms[i+1] == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				terms[i] = terms[i] / terms[i+1]
			}

			// Удаляем использованный оператор и следующее число
			terms = append(terms[:i+1], terms[i+2:]...) // Удаляем i+1
			operators = append(operators[:i], operators[i+1:]...)
			i-- // Уменьшаем i, чтобы не пропустить следующий оператор
		}
	}

	// Проверка на количество операндов после операций умножения и деления
	if len(terms) < 1 {
		return 0, fmt.Errorf("invalid number format: not enough operands after multiplication/division")
	}

	// Вычисляем результаты для + и -
	result := terms[0]
	for i := 0; i < len(operators); i++ {
		if operators[i] == "+" {
			if i+1 < len(terms) {
				result += terms[i+1]
			}
		} else if operators[i] == "-" {
			if i+1 < len(terms) {
				result -= terms[i+1]
			}
		}
	}

	return result, nil
}

// Функция для проверки, является ли символ оператором
func isOperator(ch byte) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

// Обработка выражений с плюсами и минусами
func PlusMinusExpression(result float64) (float64, error) {
	return result, nil // Можно расширить функциональность в будущем
}

// Основная функция для вычисления выражения
func Calc(s string) (float64, error) {
	if err := ValidateExpression(s); err != nil {
		return 0, err
	}

	str := s
	for first, second := FindBrackets(str); first != -1 && second != -1; first, second = FindBrackets(str) {
		if second <= first+1 {
			return 0, fmt.Errorf("invalid number format: empty brackets")
		}

		subExpr := str[first+1 : second]
		subResult, err := MulDivRes(subExpr)
		if err != nil {
			return 0, err
		}

		str = fmt.Sprintf("%s%f%s", str[:first], subResult, str[second+1:])
	}

	finalResult, err := MulDivRes(str)
	if err != nil {
		return 0, err
	}

	finalResult, err = PlusMinusExpression(finalResult)
	if err != nil {
		return 0, err
	}

	return finalResult, nil
}

// Тестирование
func main() {

	var expr string
	fmt.Scan(&expr)

	result, err := Calc(expr)
	if err != nil {
		fmt.Printf("Error calculating '%s': %s\n", expr, err)
	} else {
		fmt.Printf("Result of '%s': %f\n", expr, result)
	}

}
