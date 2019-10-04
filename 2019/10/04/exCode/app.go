package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")

	a := []int{1, 2, 3, 4, 5}
	b := 2
	fmt.Println(a)
	fmt.Println(shift(a, b))
	/*
		1 2 3 4 5  , 2

		=> 3 4 5 1 2

	*/

	// f1 := []string{"apple", "banana", "tomato"}
	// f2 := []string{"grape", "strawberry"}
	// f3 := append(f1, f2) // 이어 붙이기
}

func shift(array []int, count int) []int {

	fmt.Println("pre array ", array[:count])
	fmt.Println("post array ", array[count:])
	// preArray := array[:count]
	// postArray := array[count:]
	fmt.Println("array ", array)
	fmt.Println("count ", count)
	result := make([]int, 0, len(array))

	fmt.Println("1 result ", result)
	for _, i := range array[count:] {
		result = append(result, i)
	}
	for _, i := range array[:count] {
		result = append(result, i)
	}
	fmt.Println("2 result ", result)
	// result = append(result, preArray)
	// result = append(result, postArray)

	fmt.Println("result ", result)
	return result
}
