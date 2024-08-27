package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var ToDoList = []string{}

func main(){
	wg := new(sync.WaitGroup)
	

	for i:=1; i<100; i++{
		wg.Add(1)
		go FetchData(wg, i)
	}
	wg.Wait()
	fmt.Println(ToDoList)

}

func FetchData(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(id))
	if err != nil{
		panic("error")
	}
	responseBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil{
		panic("error")
	}
	ToDoList = append(ToDoList, string(responseBody))

}