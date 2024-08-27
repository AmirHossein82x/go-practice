package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	source1 := make(chan string)
	source2 := make(chan string)

	var result string

	go GetUrlForSpeed("https://jsonplaceholder.typicode.com/todos", source1)
	go GetUrlForSpeed("https://reqres.in/api/users", source2)

	select {

	case result = <-source1:
		fmt.Println("first site")

	case result = <-source2:
		fmt.Println("second site")

	case <-time.After(time.Second * 2):
		fmt.Println("waste of time")
	}

	fmt.Println(result)

}

func GetUrlForSpeed(url string, source chan<- string) {

	defer func() {
		err := recover()
		if err != nil {
			log.Println("error")
		}

	}()

	
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	responseBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}
	source <- string(responseBody)

}
