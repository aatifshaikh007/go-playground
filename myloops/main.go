package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops")
	days := []string{"monday", "Tuesday", "Wednesday", "Friday", "Saturaday"}
	fmt.Println(days)
	for i := 1; i < len(days); i++ {
		fmt.Println(days[i])

	}
	//if you want index also
	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}
	rougvalue := 1
	for rougvalue < 10 {
		if rougvalue == 2 {
			goto lco

		}
		fmt.Println("Value is : ", rougvalue)
		rougvalue++
	}
lco:
	fmt.Println("Jumping at LearnCodeonline.in")
}
