package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		//fmt.Println(f.Name())
		if f.Name() != "main.go" || !strings.Contains(f.Name(), "wihtChar") {
			appendColon(f.Name())
		}
	}
}
func appendColon(fileName string) {
	filepath := "./" + fileName
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	wfile, err := os.OpenFile(filepath+"wihtChar.log", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer wfile.Close()

	for scanner.Scan() {
		_, err := wfile.WriteString("" + scanner.Text() + ",\n")
		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}

	}
}
