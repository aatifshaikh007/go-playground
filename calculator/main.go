package main

import (
	"errors"
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string
	fmt.Println("welcome to the Mini calculator ")
	fmt.Println("Enter the first num : ")
	fmt.Scanln(&num1)

	fmt.Print("Enter opertor (+ - * /): ")
	fmt.Scanln(&operator)

	fmt.Println("Entere the secound number : ")
	fmt.Scanln(&num2)

	result, err := calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Printf("Result : %.2f\n", result)

}

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil

	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("cannot divide by zero ")
		}
		return a / b, nil
	default:
		return 0, errors.New("invalid operator")
	}
}
