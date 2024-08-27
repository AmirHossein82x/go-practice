package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// var results = []string{}

func main(){
	start := time.Now()
	wg := sync.WaitGroup{}
	source := make(chan string)
	for i := 1; i < 50; i++ {
		wg.Add(1)
		go GetUrl2(i, &wg, source)
	}

	for i := 1; i < 50; i++ {
		item := <- source
		fmt.Println(item)
	}
	wg.Wait()
	fmt.Println(source)
	fmt.Println(time.Since(start).Seconds())


}


func GetUrl2(item int, wg *sync.WaitGroup, source chan string){
	defer wg.Done()
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/"+ strconv.Itoa(item))
	if err != nil{
		panic(err)
	}
	responseBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil{
		panic(err)
	} 
	source <- string(responseBody)
	
}