package main 

import (
	"fmt"
)

func addNumber(num1,num2,num3 int) int {
	sum:=num1+num2+num3
	return sum
}

func main(){
	fmt.Println("Sum of Number: ",addNumber(2,3,4))
}