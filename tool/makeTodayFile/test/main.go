package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileName := "temp.txt"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		file, err := os.Create("temp.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	} else {
		currentTime := time.Now().Local()
		err = os.Chtimes(fileName, currentTime, currentTime)
		if err != nil {
			fmt.Println(err)
		}
	}
}
