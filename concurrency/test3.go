package main

import "sync"



var userList = []int{} 
func main(){
	wg := sync.WaitGroup{}
	condition := sync.Cond{L: &sync.Mutex{}}

}


func Viewer(userId int, wg *sync.WaitGroup, condition * sync.Cond){
	defer wg.Done()
	Checking(userId, condition)

}


func Checking(userId int, condition * sync.Cond){

}