package main

import "fmt"

func main() {
	mytodo := Todos{}
	mytodo.add("Aatif")
	mytodo.add("Aman")
	fmt.Println(mytodo)
	mytodo.lis()
}

type Todo struct {
	ID        int
	Title     string
	Completed bool
}
type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		ID:        len(*todos) + 1,
		Title:     title,
		Completed: false,
	}
	*todos = append(*todos, todo)

}
func (todos *Todos) lis() {
	for _, t := range *todos {
		fmt.Printf("ID : %d, title : %s, Completed : %t\n", t.ID, t.Title, t.Completed)
	}
}
func (todos *Todos) toggel(id int) {

}
