package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func main() {

	fmt.Println("test")
	out, _ := exec.Command("pwd").Output()
	fmt.Println(string(out))
	temp := strings.TrimSpace(string(out))

	fmt.Println(strings.TrimSpace(temp))
	cmd_string := "cp -r " + temp + "/dir1" + " " + temp + "/dir2"

	fmt.Println(cmd_string)

	yesterday, today := getTodayAndYesterday()
	separate := getSeparate()

	transDir := getPath(separate)

	dir1 := transDir + separate + yesterday

	dir2 := transDir + separate + today

	fmt.Println("dir1 ", dir1)
	fmt.Println("dir2 ", dir2)

	cmd := exec.Command("cp", "-r", dir1, dir2)
	stdoutStderr, _ := cmd.CombinedOutput()
	fmt.Println(string(stdoutStderr))

}

func getTodayAndYesterday() (string, string) {

	yesterdayTime := time.Now().AddDate(0, 0, -1)
	todayTime := time.Now()

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

	return yesterday, today
}

func getSeparate() string {
	separate := ""
	if runtime.GOOS == "windows" {
		separate = "\\"
	} else {
		separate = "/"
	}
	return separate
}

func getPath(separate string) string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	var ss []string

	ss = strings.Split(dir, separate)

	transDir := ""

	transDir = strings.Join(ss[0:len(ss)-3], separate)

	return transDir
}
