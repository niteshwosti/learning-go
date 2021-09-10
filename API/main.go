package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Hi there! this is my first API Practise")
}

func main(){
	http.HandleFunc("/home",handler)
	log.Fatal(http.ListenAndServe(":7000",nil))
}