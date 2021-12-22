package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("THis is main function")

	go func() {
		fmt.Println("Welcome to anonymous go routin function")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Bye to main function")
}
