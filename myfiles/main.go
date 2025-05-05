package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to the files in go")
	content := "this needss to go in file"

	file, err := os.Create("./mylocalfile.txt")
	if err != nil {
		panic(err)
	}

	lengh, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("length is ", lengh)
	defer file.Close()
	readfile("./mylocalfile.txt")
}

func readfile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	if err != nil {
		panic(err)
	}
	fmt.Println("text of the file is ", string(databyte))

}
