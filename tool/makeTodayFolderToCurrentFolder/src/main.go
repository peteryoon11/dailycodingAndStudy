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
	// dir1_high := transDir + separate

	dir2 := transDir + separate + today

	fmt.Println("!!! transDir ", transDir)
	fmt.Println("!!! dir1 ", dir1)
	fmt.Println("!!! dir2 ", dir2)
	fmt.Println("!!! yesterday ", yesterday)
	fmt.Println("!!! today ", today)

	temp_dir := ""
	for item, value := range strings.Split(today, separate) {
		fmt.Println("item!! ", item, "value ", value)
		if item == 2 {
			break
		}

		if item > 0 {
			temp_dir += separate + value
		} else {
			temp_dir += value
		}

		// cmd3 := exec.Command("mkdir", "-p", temp_dir)
		// stdoutStderr4, _ := cmd3.CombinedOutput()
		// fmt.Println(string(stdoutStderr4))
	}

	fmt.Println("temp_dir ", temp_dir)
	touch_file_name := "i_learend.md"
	cmd1 := exec.Command("mkdir", "-p", dir2)
	stdoutStderr1, _ := cmd1.CombinedOutput()
	fmt.Println(string(stdoutStderr1))

	cmd2 := exec.Command("touch", dir2+separate+touch_file_name)
	stdoutStderr2, _ := cmd2.CombinedOutput()
	fmt.Println(string(stdoutStderr2))
	// cmd := exec.Command("cp", "-r", dir1, dir2)
	cmd3 := exec.Command("cp", dir1+separate+touch_file_name, dir2+separate+touch_file_name)
	stdoutStderr3, _ := cmd3.CombinedOutput()
	fmt.Println(string(stdoutStderr3))

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
