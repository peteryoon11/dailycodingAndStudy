package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	//	fmt.Println(t.Format("yyyyMMddHHmmss"))
	fmt.Println(t)
	fmt.Println(t.AddDate(0, 1, 0))
	makeDir()
}
func makeDir() {
	os.Mkdir("./testdir", 755)
}
