package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// index , currNum
	for _, currNum := range numbers {
		if currNum%2 == 0 {
			fmt.Println(currNum, "is even")
		} else {
			fmt.Println(currNum, "is ODD")
		}
	}
}
