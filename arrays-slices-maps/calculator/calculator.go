package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string
	var num int

	fmt.Println("Mini Project")
	fmt.Print("Enter a number:  ")
	fmt.Scanln(&num1)

	fmt.Print("Entere the operator(+,-,*,/): ")
	fmt.Scanln(&operator)

	fmt.Print("Entere the secound num")
	fmt.Scanln(&num2)
	result := calculate(num1, num2, operator)
	fmt.Println("Result : ", result)
	tresult := threeadd(num1, num2, num2)
	fmt.Println(tresult)
	check := odd0reven(num)
	fmt.Println(check)
}
func calculate(a, b float64, op string) float64 {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b != 0 {
			return a / b
		}

		fmt.Println("divide By zero")
		return 0
	default:
		fmt.Println("Invalid operator")
		return 0
	}
}
func threeadd(a, b, c float64) float64 {
	return a + b + c
}
func odd0reven(num int) string {
	if num%2 == 0 {
		return "even"
	} else {
		return "odd"
	}

}
