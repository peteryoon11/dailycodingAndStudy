package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	i := checkMonthDaysNumber(time.Now())
	for i > 0 {
		i--
		fmt.Println(i)
		/* temp, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err)
		} */
		//	os.IsExist()
		tempString := strconv.Itoa(i)
		if i < 10 {
			tempString = "0" + tempString
		}
		err := os.Mkdir("./temp/"+tempString, 755)
		if err != nil {
			os.Mkdir("./temp", 755)
			fmt.Println(err)

		}
	}
	fmt.Println()
}
func checkMonthDaysNumber(checkPoint time.Time) (num int) {
	tempMonth := checkPoint.Month()
	//count := 31
	//s := make([]int, count)
	//num := 0
	//fmt.Println(tempMonth)
	for num < 31 {
		//if tempMonth == checkPoint.AddDate(0, 0, 1).Month() {
		//fmt.Println(tempMonth.String(), "   ", checkPoint.AddDate(0, 0, num).Month().String())
		if strings.EqualFold(tempMonth.String(), checkPoint.AddDate(0, 0, num).Month().String()) {
			//fmt.Print(item, " ")
			num += 1
		} else {
			return num
		}
	}

	return num
}
