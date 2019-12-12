package main

import (
	"fmt"
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

	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	fmt.Println(formatted)
}
