package main

import "fmt"

type ranNumber []int

func main() {

	oneToTen := ranNumber{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, number := range oneToTen {
		if i%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}

}
