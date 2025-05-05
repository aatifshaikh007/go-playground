package main

import "fmt"

func main() {
	fmt.Println("welcome to the Array Class")
	var fruitlist [4]string
	fruitlist[0] = "mango"
	fruitlist[1] = "graphs"
	fruitlist[3] = "banana"
	//fmt.Println(fruitlist[1])
	for i := 0; i < len(fruitlist); i++ {
		fmt.Println(fruitlist[i])
	}

}
