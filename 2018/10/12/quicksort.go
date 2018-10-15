package main

import (
	"fmt"
)

func quicksort(array []int) []int {
	if len(array) < 2 {
		return array
	} else {
		pivot := array[0]
		var less []int
		var greater []int
		for _, i := range array[1:] {

			//if n != 0 && i <= pivot {
			if i <= pivot {
				less = append(less, i)
			}
		}
		for _, i := range array[1:] {
			if i > pivot {
				greater = append(greater, i)
			}
		}
		/* var temp []int

		temp = append(temp, quicksort(less)...)
		temp = append(temp, pivot)
		temp = append(temp, quicksort(greater)...)
		*/
		//return append(quicksort(less),[pivot] , quicksort(greater)...)
		return append([]int{}, append(append(less, pivot), quicksort(greater)...)...)
		//https://stackoverflow.com/questions/16248241/concatenate-two-slices-in-go
	}
}
func main() {
	temp := []int{2, 5, 1, 10, 12, 9, 3}
	fmt.Println(quicksort(temp))

}
