package main

import (
	"fmt"
	"time"
)

/*
요구사항

1. 실행한 날짜에 맞춰서 해당 월의

한번 실행을 하면 해당 월의 일수 만큼 폴더가 생성 되고


*/
func main() {
	/* 	date := jodaTime.Format("YYYY.MM.dd", time.Now())
	   	fmt.Println(date) */
	t := time.Now()
	//	fmt.Println(t.Format("yyyyMMddHHmmss"))
	fmt.Println(t)
	fmt.Println(t.AddDate(0, 1, 0))
}
