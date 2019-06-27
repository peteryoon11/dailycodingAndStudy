package main

import (
	"fmt"

	"./polymorphism"
)

type Test struct {
	val1 int
	val2 int
}

func (type1 *Test) testFunc1() {
	fmt.Println(type1.val1, " ete")
}

func main() {
	test1 := new(Test)
	test1.val1 = 11
	test1.testFunc1()
	fmt.Println(test1.val2, "val2")

	hillary := new(polymorphism.Hillary)
	hillary.Slogan()                // "Stronger together."
	polymorphism.SaySlogan(hillary) // "Stronger together."
	trump := new(polymorphism.Trump)
	polymorphism.SaySlogan(trump) // "Make America great again."
}
