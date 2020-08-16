package main

import (
	"fmt"
	"io"
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
	yesterdayTime := time.Now().AddDate(0, 0, -1)
	todayTime := time.Now()
	fmt.Println("test")
	fmt.Println(time.Now())
	fmt.Println(time.Now().Format("2019-12-08"))

	yesterday := fmt.Sprintf("%d/%02d/%02d",
		//  T%02d:%02d:%02d",
		yesterdayTime.Year(), yesterdayTime.Month(), yesterdayTime.Day(),
		// t.Hour(), t.Minute(), t.Second()
	)

	today := fmt.Sprintf("%d/%02d/%02d",
		//  T%02d:%02d:%02d",
		todayTime.Year(), todayTime.Month(), todayTime.Day(),
		// t.Hour(), t.Minute(), t.Second()
	)

	fmt.Println(yesterday)
	fmt.Println(today)

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

	// currentDirName := ss[len(ss)-1]
	currentDirName2 := ""
	if runtime.GOOS == "windows" {
		currentDirName2 = strings.Join(ss[0:len(ss)-3], "\\")
	} else {
		currentDirName2 = strings.Join(ss[0:len(ss)-3], "/")
	}

	fmt.Println("currentDirName2 ", currentDirName2)

	dir1 := currentDirName2 + "/" + yesterday + "/i_learend.md"
	sourceFile, err := os.Open(dir1)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	dir2 := currentDirName2 + "/" + today + "/i_learend.md"
	os.MkdirAll(currentDirName2+"/"+today, 0755)
	// Create new file

	newFile, err := os.Create(dir2)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytesCopied, err := io.Copy(newFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesCopied)

	// temp, err := os.Stat(currentDirName)
	// fmt.Println(" temp ", temp)
	// fmt.Println(" err ", err)

}
/**
남은 부분 
1. repo root 에서 오늘 날짜 파일 만들기 
2. 오늘날짜 파일 만들고 git 명령어 실행해서 저장 및 메세지, 푸쉬 까지. 

*/