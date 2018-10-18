package main

import (
	"fmt"
	"strings"
)

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func search(name string, graph map[string][]string) bool {

	var search_queue = []string{}

	search_queue = append(search_queue, graph[name]...)
	var serarched = []string{}

	for len(search_queue) > 0 {
		person := search_queue[len(search_queue)-1 : len(search_queue)]
		search_queue = search_queue[:len(search_queue)-1]
		//pop
		fmt.Println(person[0] + " is a pop")
		if !Contains(serarched, person[0]) {
			fmt.Println("inner contains")
			if person_is_sellor(person[0]) {
				fmt.Println(person[0] + " is a mango seller!")
				return true
			} else {
				search_queue = append(search_queue, graph[person[0]]...)
				serarched = append(serarched, person[0])
			}
		}

	}
	return false
}
func person_is_sellor(name string) bool {

	return strings.EqualFold(name[len(name)-1:len(name)], "m")

}
func main() {

	graph := make(map[string][]string)
	graph["you"] = []string{"alice", "bob", "claire"}
	// https: //stackoverflow.com/questions/4286615/how-do-i-create-a-map16bytestring-in-go
	graph["bob"] = []string{"anju", "peggy"}
	graph["alice"] = []string{"peggy"}
	graph["claire"] = []string{"thom", "jonny"}
	graph["anju"] = []string{}
	graph["peggy"] = []string{}
	graph["thom"] = []string{}
	graph["johny"] = []string{}

	fmt.Println(search("you", graph))
}
