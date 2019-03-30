package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println(i)
		/* temp, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err)
		} */
		os.Mkdir("./"+strconv.Itoa(i), 755)
	}
	fmt.Println()
}
