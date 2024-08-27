package main

import (
	"fmt"
	"time"
)


func main(){
	source := make(chan int)
	go AddToChan4(source)
	r := <-source
	fmt.Println(r)
	time.Sleep(time.Second * 2)


}


func AddToChan4(source chan int){
	fmt.Println("before sending 1 to chanel")
	source <- 1
	fmt.Println("after sending 1 to chanel")

	fmt.Println("before sending 2 to chanel")
	source <- 2
	fmt.Println("after sending 2 to chanel")

	fmt.Println("before sending 3 to chanel")
	source <- 3
	fmt.Println("after sending 3 to chanel")
}


