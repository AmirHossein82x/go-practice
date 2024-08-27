package main

import (
	"fmt"
	"sync"
)

func main() {
	var money int = 4000000
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}
	wg.Add(1000)
	// go Task1(&wg, &money)
	// go Task2(&wg, &money)
	// go Task3(&wg, &money)
	for i := 0; i < 1000; i++ {
		go Task1(&wg, &money, &mx)
	}
	wg.Wait()
	fmt.Println(money)

}

func Task1(wg *sync.WaitGroup, money *int, mx *sync.Mutex) {
	defer wg.Done()
	mx.Lock()
	// fmt.Println("task1")
	*money -= 100
	mx.Unlock()
}

func Task2(wg *sync.WaitGroup, money *int) {
	defer wg.Done()

	fmt.Println("task2")
	*money -= 200
}

func Task3(wg *sync.WaitGroup, money *int) {
	defer wg.Done()
	fmt.Println("task3")
	*money -= 300

}
