package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	Aatif := User{"Aatif", "Aatif@gmail.com", true, 12}
	fmt.Println(Aatif.Age, Aatif.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
