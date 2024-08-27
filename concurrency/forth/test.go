package main

import (
	// "fmt"
	"time"
)

func main(){
	source := make(chan int)

	go Add(source)
	// r := <-source

	// fmt.Println(r)



}



func Add(src chan int){
	time.Sleep(time.Second * 3)
	src <- 1
}