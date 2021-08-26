package main

import (
	"fmt"
)

func checkEvenNumber(x int) bool {
	if x%2 == 0 {
		return true;
	}
	return false;
}

func main() {
	fmt.Println(checkEvenNumber(3))
}