package main

import (
	"fmt"
	// "time"
)


func main(){
	// source1 := make(chan int)
	source := make(chan int, 7)
	go AddToChan2(source)
	r := <-source
	fmt.Println(r)
	// s := <- source
	// fmt.Println(s)

	// x := <- source
	// fmt.Println(x)
	// time.Sleep(time.Second * 2)


}


func AddToChan2(source chan int){
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


