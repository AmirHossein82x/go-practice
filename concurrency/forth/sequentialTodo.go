package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	// "sync"
	"time"
)



func main(){
	start := time.Now()
	var results = []string{}
	for i := 1; i < 20; i++ {
		
		GetUrl5(i, &results)
	}
	
	fmt.Println(results)
	fmt.Println(time.Since(start).Seconds())


}


func GetUrl5(item int, results *[]string){

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