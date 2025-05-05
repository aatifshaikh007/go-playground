package main

import "fmt"

func main() {
	/*
		nums := []int{5, 10, 4}
		for i, val := range nums {
			fmt.Printf("index :%d,Value : %d\n", i, val) //normal example standard
		}
		for _, val := range nums { // If you want to only one value and that is secound one than  you use _,val for example
			fmt.Println(val)
		}
		for i := range nums {
			fmt.Println("index : ", i) // if you want to only one value and that is first that is index
		}

		arr := [3]int{1, 2, 3}
		for i, v := range arr {
			fmt.Println(v, i) // using for range with a array sama concept
		}

		m := map[string]int{"a": 1, "b": 2} // using for range with maps
		for key, value := range m {
			fmt.Println("key : ", key, "value : ", value)
		}

		str := "GO !"
		for index, char := range str {
			fmt.Printf("index: %d, char: %c\n", index, char) // using for range with string to iterate over it to geet index aand value
			//if %d is used after char that u get number of ascii code and use %c than get character
		}*/
	nums := []int{10, 20, 30}

	fmt.Println("Before range")
	for i := range nums {
		fmt.Println("nums[%d] address: %p, value: %d\n", i, &nums[i], nums[i])
	}
	fmt.Println("\nUsing range with value (val):")
	for _, val := range nums {
		fmt.Printf("val address: %p, value: %d\n", &val, val)
		val = val * 2 // This will NOT modify the original nums
	}
	fmt.Println("\nAfter range:")
	for i := range nums {
		fmt.Printf("nums[%d] address: %p, value: %d\n", i, &nums[i], nums[i])
	}
	fmt.Println("\nNow modifying using index:")
	for i := range nums {
		nums[i] *= 2 // This will modify the original nums
	}
	fmt.Println("\nAfter modifying by index:")
	for i := range nums {
		fmt.Printf("nums[%d] address: %p, value: %d\n", i, &nums[i], nums[i])
	}

}
