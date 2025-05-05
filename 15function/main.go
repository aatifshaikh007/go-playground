package main

import "fmt"

func main() {
	greeter()
	result := adder(5, 6)
	fmt.Println("Result is :", result)
	proresult, promessage := proadder(5, 2, 6, 3, 5, 2)
	fmt.Println("the result is", proresult)
	fmt.Println("the Pro message : ", promessage)
}
func greeter() {
	fmt.Println("hello from gplang")
}
func adder(valone int, valtwo int) int {
	return valone + valtwo
}
func proadder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "hi pro result fuction"
}
