package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

/**
목표 :
1 . 실행시에 오늘 날짜를 받아온다.
2. 받아온 오늘 날짜를 기반으로 어제 날짜 폴더를 복사해서 오늘날짜 폴더를 만든다.
*/

func main() {
	/* 	t := time.Now()
	   	//time.Now().f
	   	fmt.Println(t.Format("2019-02-01 00:00:00:000"))
	   	fmt.Println(t)
	   	fmt.Println(t.AddDate(0, 1, 0))
	   	fmt.Println(t.AddDate(0, 1, 0).Month()) */
	//fmt.Println(t.AddDate(0, 1, 0).Sub(t.da))
	//fmt.Println(t.Month())
	/* 	fmt.Println(t.Day())
	   	fmt.Println(t.AddDate(0, 1, 0).Day()) */
	// 오늘부터 다음달 전까지 며칠 남았나?
	/*test := 20
	 s := make([]int, test)
	for item := range s {
		fmt.Println(item)
	} */
	//makeDir()
	//	testFor()
	//dateFormat()
	fmt.Print(checkMonthDaysNumber(time.Now()))
	fmt.Println(" remains days")
}
func checkMonthDaysNumber(checkPoint time.Time) (num int) {
	tempMonth := checkPoint.Month()
	//count := 31
	//s := make([]int, count)
	//num := 0
	//fmt.Println(tempMonth)
	for num < 31 {
		//if tempMonth == checkPoint.AddDate(0, 0, 1).Month() {
		fmt.Println(tempMonth.String(), "   ", checkPoint.AddDate(0, 0, num).Month().String())
		if strings.EqualFold(tempMonth.String(), checkPoint.AddDate(0, 0, num).Month().String()) {
			//fmt.Print(item, " ")
			num += 1
		} else {
			return num
		}
	}

	return num
}
func makeDir(tempdir string) {
	err := os.Mkdir("./"+tempdir, 755)
	if err != nil {
		log.Println(err)
		//fmt.Println(err)
	}
}

func testFor() {
	nums := []int{1, 2, 3}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum : ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index : ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func dateFormat() {
	p := fmt.Println

	// Here's a basic example of formatting a time
	// according to RFC3339, using the corresponding layout
	// constant.
	t := time.Now()
	p(t.Format(time.RFC3339))

	// Time parsing uses the same layout values as `Format`.
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// `Format` and `Parse` use example-based layouts. Usually
	// you'll use a constant from `time` for these layouts, but
	// you can also supply custom layouts. Layouts must use the
	// reference time `Mon Jan 2 15:04:05 MST 2006` to show the
	// pattern with which to format/parse a given time/string.
	// The example time must be exactly as shown: the year 2006,
	// 15 for the hour, Monday for the day of the week, etc.
	p(t.Format("3:04PM"), " =========")
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// For purely numeric representations you can also
	// use standard string formatting with the extracted
	// components of the time value.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// `Parse` will return an error on malformed input
	// explaining the parsing problem.
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
