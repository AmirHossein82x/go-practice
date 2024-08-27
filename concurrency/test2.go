package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(4)
	go Sum(&wg)
	go Sum(&wg)
	go Sum(&wg)
	go Sum(&wg)
	go Sum(&wg)
	go Sum(&wg)
	go Sum(&wg)
	go Sum(&wg)
	
	wg.Wait()
	fmt.Println(time.Since(start).Seconds())
}


func Sum(wg *sync.WaitGroup){
	a:=1
	defer wg.Done()
	for i := 0; i < 2000000000; i++ {
		a += i
	}
	 
}