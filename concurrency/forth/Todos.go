package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"
)



func main(){
	start := time.Now()
	var results = []string{}
	wg := sync.WaitGroup{}
	for i := 1; i < 50; i++ {
		wg.Add(1)
		go GetUrl(i, &wg, &results)
	}
	wg.Wait()
	fmt.Println(results)
	fmt.Println(time.Since(start).Seconds())


}


func GetUrl(item int, wg *sync.WaitGroup, results *[]string){
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
	*results = append(*results, string(responseBody))
}