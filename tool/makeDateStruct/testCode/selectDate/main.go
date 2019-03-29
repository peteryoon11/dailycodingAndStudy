package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	//	fmt.Println(t.Format("yyyyMMddHHmmss"))
	fmt.Println(t)
	fmt.Println(t.AddDate(0, 1, 0))
}
