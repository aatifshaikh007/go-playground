package main

import "fmt"

func main() {
	fmt.Println("slice class")
	var fruitlist = []string{"apple", "Tomato", "Banana"}
	fmt.Println(fruitlist)
	fruitlist = append(fruitlist, "Mango", "Banana")
	fmt.Println(fruitlist)
	var index int = 1
	fruitlist = append(fruitlist[:index], fruitlist[index+1:]...)
	fmt.Println(fruitlist)
}
