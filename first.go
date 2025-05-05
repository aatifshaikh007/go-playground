package main

import "fmt"

func main() {
	var name string = "Aatif"
	age := 15
	mounth := "junz"
	const country = "India"
	isGopher := true
	var pi float64 = 3.1415
	day := "Monday"
	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("Country:", country)
	fmt.Println("pi value:", pi)
	fmt.Println("is gopher:", isGopher)
	sum := Add(10, 5)
	fmt.Println(sum)
	quo, rem := divide(10, 2)
	fmt.Println("quo", quo)
	fmt.Println("rem", rem)

	if age < 18 {
		fmt.Println("minor")
	} else if age == 18 {
		fmt.Println("Just became an adult")
	} else {
		fmt.Println("Adult")
	}
	switch day {
	case "Monday":
		fmt.Println("Start of the week ")
		fallthrough

	case "Friday":
		fmt.Println("Almost Weekend !")
	default:
		fmt.Println("midweek")
	}

	switch mounth {
	case "jan", "feb", "march":
		fmt.Println("winter")
	case "jun", "july", "sep":
		fmt.Println("rainy")
	}

}
func Add(x int, y int) int {
	result := x + y
	//fmt.Println("sum ", result)
	return result
}
func divide(a int, b int) (int, int) {
	quo := a / b
	rem := a % b
	return quo, (rem)
}
