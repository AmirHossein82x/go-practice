// package main

// import (
// 	"fmt"
// 	"sync"
// )


// var money = 400
// func main(){
// 	// wg := sync.WaitGroup{}
	// mx := sync.Mutex{}
// 	for i := 0; i < 3; i++ {
// 		go Cost(&money, &mx)
// 	}
// 	fmt.Println(money)

// }

// func Cost(money *int, mx *sync.Mutex) {
// 	mx.Lock()
// 	*money -= 100
// 	mx.Unlock()
	
// }


