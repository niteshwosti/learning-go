package main

import (
	"fmt"
)
	
func addNumber(num1 int, num2 int) int {
	return num1+num2
}

func main() {
sum:= addNumber(10,5)
fmt.Println(sum)
}