package main

import (
	"fmt"
)

func checkEvenNumber(x int) {
	if x%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}
	
}

func main() {
	checkEvenNumber(3)
}