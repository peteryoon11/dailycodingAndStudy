package main

import "fmt"

func main() {

	pibonachi := make([]int, 0)
	for i := 0; i < 10; i++ {
		if i < 2 {
			pibonachi = append(pibonachi, i)
			/*} else if i == 1 {
				pibonachi = append(pibonachi, 1)
			} */
		} else {
			pibonachi = append(pibonachi, pibonachi[i-1]+pibonachi[i-2])
		}
	}
	fmt.Println(pibonachi)
}
