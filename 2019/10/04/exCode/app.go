package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")

	a := []int{1, 2, 3, 4, 5}
	b := 2
	fmt.Println(a)
	shift(a, b)
}

func shift(array []int, count int) {
	fmt.Println("array ", array)
	fmt.Println("count ", count)
}
