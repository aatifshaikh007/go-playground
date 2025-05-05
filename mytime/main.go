package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time ")
	presenttime := time.Now()
	fmt.Println(presenttime.Format("01-02-2006 15:04:05 Monday"))

}
