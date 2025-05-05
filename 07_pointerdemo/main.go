package main

import "fmt"

func main() {
	var a int = 42
	var p *int = &a
	fmt.Println("valiue of a : ", a)
	fmt.Println("Address of a : ", &a)
	fmt.Println("value of p (address) : ", p)
	fmt.Println("Value at address p (*p):", *p)
	*p = *p + 10
	fmt.Println("value of a ", a)
}
