package main

import (
	"fmt"
	"sync"
)


func main(){ // simple
	mx := new(sync.Mutex)
	wg := new(sync.WaitGroup)

	counter := 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mx.Lock()
			counter += 1
			mx.Unlock()

		}()
	}
	wg.Wait()
	fmt.Println(counter)
}