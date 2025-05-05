package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	Aatif := User{"Aatif", "Aatif@gmail.com", true, 12}
	fmt.Println(Aatif.Age, Aatif.Email)
	Aatif.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("IS user active ", u.Status)
}
