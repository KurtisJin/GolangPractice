package main

import "fmt"

type oddOrEven []int
type trueOrFalse bool

func main() {

}

func oddOrEvenNumber() trueOrFalse {

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := range numbers {
		if i%2 == 0 {
			return true
		} else {
			return false
		}
	}

}
