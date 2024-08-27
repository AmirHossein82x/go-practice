package main

import "fmt"


func main(){
	source := make(chan int)
	go AddToChan(source, 6)
	receiver := <-source
	fmt.Println(receiver)



}


func AddToChan(source chan int, element int){
	source <- element
}


