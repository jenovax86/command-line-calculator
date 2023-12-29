package main

import (
	"fmt"
)

func calculater() int {
	var firstNumber int
	var secondNumber int
	var result int
	var operator string

	fmt.Scan(&firstNumber, &operator, &secondNumber)

	switch operator {
	case "+":
		result = firstNumber + secondNumber
		break

	case "-":
		result = firstNumber - secondNumber
		break
	case "*":
		result = firstNumber * secondNumber
		break
	case "/":
		result = firstNumber / secondNumber
		break
	default:
		fmt.Println("default")
	}

	return result
}

func main() {
	calculate := calculater()
	fmt.Println(calculate)
}
