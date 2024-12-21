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

	t.Run("Сложение", func(t *testing.T) {
		assertResult(t, "2 + 3", 5, false)
	})

	t.Run("Вычитание", func(t *testing.T) {
		assertResult(t, "10 - 4", 6, false)
	})

	t.Run("Умножение", func(t *testing.T) {
		assertResult(t, "3 * 4", 12, false)
	})

	t.Run("Деление", func(t *testing.T) {
		assertResult(t, "8 / 2", 4, false)
	})

	t.Run("Все доступные операции", func(t *testing.T) {
		assertResult(t, "2 + 3 * (4 - 1)", 11, false)
	})

	t.Run("Деление на ноль", func(t *testing.T) {
		assertResult(t, "5 / 0", 0, true)
	})

	t.Run("Выражение с пробелами", func(t *testing.T) {
		assertResult(t, "  4 +  5 * 2 ", 14, false)
	})

	t.Run("Отрицательное число", func(t *testing.T) {
		assertResult(t, "-5 + 3", -2, false)
	})

	t.Run("Несколько скобок", func(t *testing.T) {
		assertResult(t, "((2 + 3) * (4 + 1))", 25, false)
	})

	t.Run("Неподдерживаемые символы", func(t *testing.T) {
		assertResult(t, "2 + a", 0, true)
	})

	t.Run("Нет закрывающей скобки", func(t *testing.T) {
		assertResult(t, "(2 + 3", 0, true)
	})

	t.Run("Пустой ввод", func(t *testing.T) {
		assertResult(t, "", 0, true)
	})

	t.Run("Двойной плюс", func(t *testing.T) {
		assertResult(t, "2 ++ 2", 0, true)
	})
}
