package main 

import(
	"fmt"
)
func main(){
	var x,y int = 1,2 
	var f float64=float64(x*y)
	var z uint = uint(f)
	fmt.Println(x,y,z)
}