package main 

import "fmt"

type Demo struct{
	X int 
	Y int
}
func main(){
	fmt.Println(Demo{1,2})
}