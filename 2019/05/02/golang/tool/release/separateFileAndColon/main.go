package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var loopStandard = 1000

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	/* 	for _, f := range files {
		if f.Name() != "main.go" && !strings.Contains(f.Name(), "wihtChar") && f.Name() != "origin" {
			file, err := os.Open("./" + f.Name())
			if err != nil {
				log.Fatal(err)
			}
			//fmt.Println(lineCounter(file))
			num, err := lineCounter(file)
			if err != nil {
				log.Fatal(err)
			}

			if num < loopStandard {
				appendColon(f.Name(), 1)
			} else {
				loopCount := num / loopStandard
				for loopCount > 0 {
					fmt.Println("loopcount  ", loopCount)
					appendColon(f.Name(), loopCount)
					loopCount--
				}
			}

		}

	} */
	/* for _, f := range files {
		//fmt.Println(f.Name())
		if f.Name() != "main.go" && !strings.Contains(f.Name(), "wihtChar") && f.Name() != "origin" {

			appendColon(f.Name())
		}
	} */
	for _, f := range files {
		//fmt.Println(f.Name())
		if f.Name() != "main.go" && !strings.Contains(f.Name(), "wihtChar") && f.Name() != "origin" {
			//	fmt.Println("f name  ", f.Name())
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
	num, err := lineCounter(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file, err = os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	loopCount := 0

	tempArray := make([]string, 0)
	for scanner.Scan() {
		//fmt.Println("tempArray  ", len(tempArray))
		//fmt.Println(scanner.Text())
		tempArray = append(tempArray, scanner.Text())
	}

	fmt.Println("tempArray len ", len(tempArray))

	if num < loopStandard {
		//	appendColon(f.Name(), 1)
		makeFileAndSaveLine(filepath, 1, tempArray)
	} else {
		loopCount = num / loopStandard
		count := 1
		loopCount++
		fmt.Println("loopCount ", loopCount)
		for loopCount >= count {
			/*
				0, 999
				1000 , 1999
				2000 , 2999
			*/
			//makeFileAndSaveLine(filepath, loopCount, tempArray[ (count-1)*loopStandard:count*loopStandard-1])
			if count*loopStandard-1 > len(tempArray) {

				//	fmt.Println("count ", count, " ", (count-1)*loopStandard, " ", len(tempArray))
				makeFileAndSaveLine(filepath, count, tempArray[(count-1)*loopStandard:len(tempArray)])
			} else {
				//fmt.Println("count ", count, " ", (count-1)*loopStandard, " ", count*loopStandard)
				makeFileAndSaveLine(filepath, count, tempArray[(count-1)*loopStandard:count*loopStandard])
			}
			//	fmt.Println("count ", count, " ", (count-1)*loopStandard, " ", count*loopStandard-1)
			count++

		}
	}
	//wfile, err := os.OpenFile(filepath+strconv.Itoa(num)+"wihtChar.log", os.O_WRONLY|os.O_CREATE, 0644)
	/* wfile, err := os.OpenFile(filepath+strconv.Itoa(num)+"wihtChar.log", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	} */
	/* 	wfile, err := os.OpenFile(filepath+strconv.Itoa(loopCount)+"wihtChar.log", os.O_WRONLY|os.O_CREATE, 0644)
	   	if err != nil {
	   		log.Fatalf("failed opening file: %s", err)
	   	}

	   	for scanner.Scan() {

	   		_, err = wfile.WriteString("'" + scanner.Text() + "',\n")
	   		if err != nil {
	   			log.Fatalf("failed writing to file: %s", err)
	   		}

	   		realCount++
	   	}

	   	defer wfile.Close() */
}

func makeFileAndSaveLine(filepath string, num int, lineArray []string) {
	if len(lineArray) > 1001 {
		fmt.Println("can make 1000 line ", len(lineArray))
		return
	}
	wfile, err := os.OpenFile(filepath+strconv.Itoa(num)+"wihtChar.log", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	for _, content := range lineArray {
		//	fmt.Println("content ", content)
		_, err = wfile.WriteString("'" + content + "',\n")
		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}
	}
	defer wfile.Close()
}

// https://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently
func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
