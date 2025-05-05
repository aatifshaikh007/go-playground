package main

import "fmt"

func main() {
	fmt.Println("map  in golang ")

	lang := make(map[string]string)
	lang["js"] = "Javascript"
	lang["RB"] = "Ruby"
	lang["jv"] = "Java"
	fmt.Println(lang["RB"])

	for _, value := range lang {
		fmt.Println(value)
	}

}
