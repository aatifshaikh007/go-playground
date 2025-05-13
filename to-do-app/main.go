package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	mytodo := Todos{}
	mytodo.add("Aatif")
	mytodo.add("Aman")
	fmt.Println(mytodo)
	mytodo.lis()
}

type Todo struct {
	ID             int
	Title          string
	Completed      bool
	completiontime *time.Time
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

func (todos *Todos) validates(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (todos *Todos) toggle(index int) error {
	if err := todos.validates(index); err != nil {
		return err
	}

	todo := (*todos)
	iscompleted := todo[index].Completed
	if !iscompleted {
		completiontime := time.Now()
		todo[index].completiontime = &completiontime
	}

	todo[index].Completed = !iscompleted
	return nil

}

func (todos *Todos) delete(index int) error {
	if err := todos.validates(index); err != nil {
		return err
	}

	todo := *todos
	*todos = append(todo[:index], todo[index+1:]...)
	return nil
}
