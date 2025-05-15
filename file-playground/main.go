package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//file, err := os.Create("first.txt")

	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()

	//fmt.Fprintf(file, "hello world")

	B, err := os.Open("first.txt")

	if err != nil {
		fmt.Print(err)
	}

	defer B.Close()

	content, err := io.ReadAll(B)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

}

func saveline() {
	newfile, _ := os.Create("new.txt")
	defer newfile.Close()

	line := []string{"line 1", "line 2", "line 3"}

	totalBytes := 0
	for _, line := range line {
		bytes, _ := fmt.Fprintln(newfile, line)
		totalBytes += bytes
	}
	fmt.Printf("%d", totalBytes)

}
