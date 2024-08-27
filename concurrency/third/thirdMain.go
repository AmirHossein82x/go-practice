package main

import (
	"fmt"
	"sync"
	"time"
	// "time"
)

var readers = 0
var canRead = false

func main(){
	condition := sync.Cond{L: &sync.Mutex{}}
	wg := new(sync.WaitGroup)
	go Reader(3, &condition, wg) 
	go Reader(6, &condition, wg) 
	go Reader(8, &condition, wg) 

	for i := 0; i < 20; i++ {
		Writer(i,&condition)
	}

	
	wg.Wait()
	
}


func Reader(Id int, condition *sync.Cond, wg *sync.WaitGroup){
	fmt.Println(readers)
	condition.L.Lock()
	defer condition.L.Unlock()
	if !canRead{
		condition.Wait()
	}
	fmt.Printf("reader %d is reading\n", Id)

}

func Writer(Id int, condition *sync.Cond){
	condition.L.Lock()
	defer condition.L.Unlock()
	fmt.Printf("writer %d is writing\n", Id)
	time.Sleep(time.Second)
	readers += 1
	if readers > 10{
		canRead = true
		condition.Broadcast()
	}
	

}