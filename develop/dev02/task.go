package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	ErrIncorrectString = errors.New("incorrect string")
)

func main() {
	var s string

	if _, err := fmt.Scan(&s); err != nil {
		log.Fatal(err)
	}

	result, err := Unpack(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func Unpack(s string) (string, error) {
	runes := []rune(s) // конечно же, работать надо с unicode
	if len(runes) == 0 {
		return "", nil
	}

	var (
		result strings.Builder // Билдер для итоговой строки
		prev   rune            // Предыдущий символ (для умножения числа)
		escape bool            // Флаг для escape последовательности
	)

	// Логика: итерируемся по символам unicode
	for i := 0; i < len(runes); i++ {
		switch {
		case runes[i] == '\\' && !escape: // меняем флаг, если текущий символ \ и до него флаг не был true (проверка для кейса \\)
			escape = true

		case escape: // кейс для символа после экранирования
			result.WriteRune(runes[i])
			prev = runes[i]
			escape = false

		case unicode.IsDigit(runes[i]): // если текущий символ - цифра
			if prev == 0 { // если до цифры ничего не было (в т.ч экранирования) - возвращаем ошибку
				return "", ErrIncorrectString
			}

			// создаем подстроку, потому что цифр может быть несколько (число > 9), в которую будем пушить цифры
			var temp string

			// Тут я решил сделать К Р А С И В О.
			// Идем доп циклом потому что нам надо сформировать подстроку temp (цифр может быть несколько, следовательно, итераций тоже)
			// Однако заново проходить эти символы в основном цикле смысла не имеет, поэтому j должен поменять i
			for j := &i; *j < len(runes); *j++ {
				if !unicode.IsDigit(runes[*j]) {
					*j-- // важно вернуться на прошлый символ, потому что если этого не делать, то основной цикл пропустит его
					break
				}
				// формируем подстроку до тех пор, пока не встретится символ !цифра
				temp += string(runes[*j])
			}
			count, _ := strconv.Atoi(temp)
			result.WriteString(strings.Repeat(string(prev), count-1)) // count-1 потому что prev уже был записан 1 раз

		default: // если текущий символ "обычный", то просто записываем
			result.WriteRune(runes[i])
			prev = runes[i]
		}
	}

	if escape { // если строка оканчивается на \, то эта строка неверная
		return "", ErrIncorrectString
	}
	return result.String(), nil
}
