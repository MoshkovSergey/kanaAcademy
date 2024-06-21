package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romeToArabic = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

// arabicToRome converts an arabic number to its roman representation
func arabicToRome(num int) string {
	if num <= 0 || num > 3999 {
		return "Недопустимое число для конвертации"
	}

	// Definition of Roman number symbols and their respective meanings
	romanNumerals := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	// Passing through the symbol map and converting to Roman numbers
	result := ""
	for _, value := range []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1} {
		for num >= value {
			result += romanNumerals[value]
			num -= value
		}
	}

	return result
}

// isRoman determines if the input string is a roman numeral
func isRoman(input string) bool {
	_, ok := romeToArabic[input]
	return ok
}

// toArabic converts a roman numeral to its arabic representation
func toArabic(roman string) int {
	return romeToArabic[roman]
}

// toRoman converts a roman numeral to its arabic representation
func toRoman(arabic int) string {
	if arabic <= 0 {
		panic("Невозможно представить результат меньше или равный 0 в римских цифрах")
	}
	return arabicToRome(arabic)
}

// calculate performs the operation specified by the given operator on the two given operands
func calculate(operand1 int, operand2 int, operator string) int {
	if operand1 <= 0 || operand1 > 10 {
		panic("Невозможно представить результат для числа вне диапазона от 1 до 10")
	} else if operand2 <= 0 || operand2 > 10 {
		panic("Невозможно представить результат для числа вне диапазона от 1 до 10")
	}

	switch operator {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		if operand2 == 0 {
			panic("на ноль делить нельзя")
		}
		return operand1 / operand2
	default:
		panic("неизвестный оператор: " + operator)
	}
}

func main() {
	fmt.Print("Введите выражение в формате (число оператор число), например, '2 + 3': ")
	reader := bufio.NewReader(os.Stdin)
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)
	parts := strings.Split(expression, " ")
	
	if len(parts) != 3 {
		fmt.Println("Неверный формат выражения")
		return
	}

	containsRoman := isRoman(parts[0]) || isRoman(parts[2])
	containsArabic := !isRoman(parts[0]) || !isRoman(parts[2])

	if containsRoman && containsArabic {
		panic("В выражении присутствуют и римские, и арабские цифры")
	}

	var operand1, operand2 int
	var operator string
	var err error

	if containsRoman {
		operand1 = toArabic(parts[0])
		operand2 = toArabic(parts[2])
	} else {
		operand1, err = strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Неверный формат числа")
			return
		}

		operand2, err = strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("Неверный формат числа")
			return
		}
	}

	operator = parts[1]

	result := calculate(operand1, operand2, operator)

	if containsRoman {
		fmt.Printf("Результат: %s\n", toRoman(result))
	} else {
		fmt.Printf("Результат: %d\n", result)
	}
}