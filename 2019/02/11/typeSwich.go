package main

import (
	"fmt"
	"reflect"
)

type RequestObject struct {
	Test1 string `json:"Test1"`
	Test2 string `json:"Test2"`
}

func main() {

	typeSwitch(RequestObject{"test1", "test2"})
}

func typeSwitch(value interface{}) {
	switch value.(type) {
	case RequestObject:
		fmt.Println("type:", reflect.ValueOf(value).Type())
		fmt.Println(value, "is CoinProvsionRequestObject.")
	case int:
		fmt.Println("type:", reflect.ValueOf(value).Type())
		fmt.Println(value, "is int.")
	case string:
		fmt.Println("type:", reflect.ValueOf(value).Type())
		fmt.Println(value, "is string.")
	default:
		fmt.Println("type:", reflect.ValueOf(value).Type())
	}
}
