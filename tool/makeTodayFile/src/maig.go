package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
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
	showTodayDate()
}
func showTodayDate() {
	t := time.Now()
	fmt.Println("test")
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("2019-12-08"))

	formatted := fmt.Sprintf("%d/%02d/%02d",
		//  T%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		// t.Hour(), t.Minute(), t.Second()
	)

	fmt.Println(formatted)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	var ss []string
	if runtime.GOOS == "windows" {
		ss = strings.Split(dir, "\\")
	} else {
		ss = strings.Split(dir, "/")
	}

	currentDirName := ss[len(ss)-1]
	currentDirName2 := ""
	if runtime.GOOS == "windows" {
		currentDirName2 = strings.Join(ss[0:len(ss)-1], "\\")
	} else {
		currentDirName2 = strings.Join(ss[0:len(ss)-1], "/")
	}

	fmt.Println("currentDirName2 ", currentDirName2)

	fmt.Println("len ", len(ss))
	for item, value := range ss {
		fmt.Println(item, " ", value)
	}

	fmt.Println("Current Directory Name: ", currentDirName)
	fmt.Println(os.Stat(currentDirName))

	temp, err := os.Stat(currentDirName)
	fmt.Println(" temp ", temp)
	fmt.Println(" err ", err)

}
