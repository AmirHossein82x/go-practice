package main

import (
	"sync"
	"fmt"

)



var userList []int = []int{}
var ready = false

func main(){
	condition := sync.Cond{L: &sync.Mutex{}}
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go NewRequest(i, &condition, &wg)
	}
	if !ready{
		wg.Wait()

	}
	fmt.Printf("%+v", userList)

}

func NewRequest(userId int, condition *sync.Cond, wg *sync.WaitGroup){
	Checking(userId, condition, wg)
	condition.L.Lock()
	defer condition.L.Unlock()
	for !ready{
		condition.Wait()
	}
	fmt.Println("streaming  starts")

}


func Checking(userId int, condition *sync.Cond, wg *sync.WaitGroup){
	fmt.Println(userId, " is waiting for streaming")
	defer wg.Done()
	condition.L.Lock()
	defer condition.L.Unlock()
	userList = append(userList, userId)
	if len(userList) == 55{
		ready = true
		condition.Broadcast()
	}
	
}
