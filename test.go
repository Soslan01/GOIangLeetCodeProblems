package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Констатные римские цифры
const (
	I = 1
	V = 5
	X = 10
)

// Конвертация арабского числа в римское
func arabicToRoman(num int) string {
	values := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	for i, value := range values {
		for num >= value {
			num -= value
			roman += numerals[i]
		}
	}
	return roman
}

// Конвертация римского числа в арабское
func romanToArabic(roman string) (int, error) {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	arabic := 0
	for i, r := range roman {
		value, ok := romanMap[r]
		if !ok {
			return 0, errors.New("Неверная")
		}

		if i+1 < len(roman) {
			nextValue, _ := romanMap[rune(roman[i+1])]
			if value < nextValue {
				arabic -= value
			} else {
				arabic += value
			}
		} else {
			arabic += value
		}
	}

	return arabic, nil
}

// Функция для выполнения математических операций
func calculate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("Число не может быть нулевым")
		}
		return a / b, nil
	default:
		return 0, errors.New("Неверная математическая операция")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите математическу операцию :")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка считывания введенных данных:", err)
		return
	}

	input = strings.TrimSpace(input)
	tokens := strings.Split(input, " ")
	if len(tokens) != 3 {
		fmt.Println("Неверный формат ввода")
		return
	}

	var a, b int
	var romanA, romanB bool

	a, err = strconv.Atoi(tokens[0])
	if err != nil {
		a, err = romanToArabic(tokens[0])
		if err != nil {
			fmt.Println("Ошибка конветирования чисел:", err)
			return
		}
		romanA = true
	}

	b, err = strconv.Atoi(tokens[2])
	if err != nil {
		b, err = romanToArabic(tokens[2])
		if err != nil {
			fmt.Println("Ошибка конвертирования чисел:", err)
			return
		}
		romanB = true
	}

	if romanA != romanB {
		fmt.Println("Введенные числа должны быть одной системы исчисления")
		return
	}

	result, err := calculate(a, b, tokens[1])
	if err != nil {
		fmt.Println("Неверная математическая операция:", err)
		return
	}

	if romanA {
		if result <= 0 {
			fmt.Println("Числа не могут быт отрицательными")
			return
		}
		fmt.Println("Result:", arabicToRoman(result))
	} else {
		fmt.Println("Result:", result)
	}
}
