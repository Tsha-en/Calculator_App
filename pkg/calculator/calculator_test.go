package calculator_test

import (
	"testing"

	"github.com/Tsha-en/Calculator_App/pkg/calculator"
)

func assertResult(t *testing.T, input string, expected float64, errExpected bool) {
	result, err := calculator.Calc(input)

	if errExpected {
		if err == nil {
			t.Errorf("Ожидалась ошибка для выражения '%s', но ошибка не получена", input)
		}
	} else {
		if err != nil {
			t.Errorf("Не ожидалась ошибка для выражения '%s', но получена ошибка: %v", input, err)
		} else if result != expected {
			t.Errorf("Для выражения '%s' ожидался результат %v, но получен %v", input, expected, result)
		}
	}
}

func TestCalc(t *testing.T) {

	t.Run("Addition", func(t *testing.T) {
		assertResult(t, "2 + 3", 5, false)
	})

	t.Run("Subtraction", func(t *testing.T) {
		assertResult(t, "10 - 4", 6, false)
	})

	t.Run("Multiplication", func(t *testing.T) {
		assertResult(t, "3 * 4", 12, false)
	})

	t.Run("Division", func(t *testing.T) {
		assertResult(t, "8 / 2", 4, false)
	})

	t.Run("Complex expression", func(t *testing.T) {
		assertResult(t, "2 + 3 * (4 - 1)", 11, false)
	})

	t.Run("Division by zero", func(t *testing.T) {
		assertResult(t, "5 / 0", 0, true)
	})

	t.Run("Expression with spaces", func(t *testing.T) {
		assertResult(t, "  4 +  5 * 2 ", 14, false)
	})

	t.Run("Negative numbers", func(t *testing.T) {
		assertResult(t, "-5 + 3", -2, false)
	})

	t.Run("Nested parentheses", func(t *testing.T) {
		assertResult(t, "((2 + 3) * (4 + 1))", 25, false)
	})

	// Тесты для некорректных выражений
	t.Run("Invalid characters", func(t *testing.T) {
		assertResult(t, "2 + a", 0, true)
	})

	t.Run("Mismatched parentheses", func(t *testing.T) {
		assertResult(t, "(2 + 3", 0, true)
	})

	t.Run("Empty input", func(t *testing.T) {
		assertResult(t, "", 0, true)
	})

	t.Run("Extra operators", func(t *testing.T) {
		assertResult(t, "2 ++ 2", 0, true)
	})
}
