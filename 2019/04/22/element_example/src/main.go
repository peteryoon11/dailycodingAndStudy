package main

import (
	"fmt"
)

func main() {
	/*
	*
	**
	***
	****
	*****
	******
	*******
	********
	*********
	**********

	 */
	for i := 0; i < 10; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()

	}
	fmt.Println("\n\n\n")
	/*

	**********
	*********
	********
	*******
	******
	*****
	****
	***
	**
	*

	 */
	for i := 0; i < 10; i++ {
		for j := 10; j > i; j-- {
			fmt.Print("*")
		}
		fmt.Println()

	}

	fmt.Println("\n\n\n")
	/*
			 **
		    ****
		   ******
		  ********
		 **********


	*/
	for i := 0; i < 5; i++ {
		for k := i; k < 5; k++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("**")
		}
		/* 	for j := 0; j < i; j++ {
			fmt.Print("*")
		} */
		fmt.Println()

	}

	fmt.Println("\n\n\n")
	/*
		**********
		 ********
		  ******
		   ****
		    **


	*/
	for i := 0; i < 5; i++ {
		for k := 0; k < i; k++ {
			fmt.Print(" ")
		}
		for j := 5; j > i; j-- {
			fmt.Print("**")
		}
		/* 	for j := 0; j < i; j++ {
			fmt.Print("*")
		} */
		fmt.Println()

	}
	fmt.Println("\n\n\n")
	/*
			 **
		    ****
		   ******
		  ********
		 **********


	*/
	for i := 0; i < 5; i++ {
		for k := i; k < 5; k++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			if j == 0 && i == 0 {
				fmt.Print("**")
			} else {
				/* for k := j; j <= k; k++ {

				} */

				//temp := ""
				//temp.append
				//var buffer bytes.Buffer
				//	for

				for num := 0; num <= j; num++ {

					//	buffer.WriteString("**")

					if j == 0 && num/2 == 0 {
						fmt.Print("*")
						//} else if num/2 == 0 {
						//	fmt.Print("*")
					} else {
						fmt.Print("-")
					}

					if num == j {
						fmt.Print("*")
					}
				}

			}

		}
		/* 	for j := 0; j < i; j++ {
			fmt.Print("*")
		} */
		fmt.Println()

	}

}
