package main

import "fmt"

func main() {
	i, j := 10, 20

	p := &i         
	fmt.Println(*p) 
	*p = 21         
	fmt.Println(i)  

	p = &j         
	*p = *p / 5  
	fmt.Println(j) 
}
