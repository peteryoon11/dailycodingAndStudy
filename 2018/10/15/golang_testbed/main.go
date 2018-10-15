package main

import (
	"fmt"
)

func main() {
	var search_queue = []string{}
	search_queue = append(search_queue, "test")
	search_queue = append(search_queue, "test2")
	for _, item := range search_queue {
		fmt.Println(item)
	}
	temp := search_queue[len(search_queue)-1 : len(search_queue)]
	fmt.Println("\n\n")
	for _, item := range temp {
		fmt.Println(item)
	}
}
