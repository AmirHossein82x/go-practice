package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Todo struct {
	UserId    int `json:"userId"`
	Id        int `json:"id"`
	Title     string `json:"title"`
	Completed bool `json:"completed"`
}

func main() {
	start := time.Now()
	source := make(chan []byte)
	todo := Todo{}
	todos := []Todo{}
	for i := 1; i < 41; i++ {
		go GetUrl(i, source)
	}
	for j := 1; j < 41; j++ {
		json.Unmarshal(<-source, &todo)
		todos = append(todos, todo)

	}

	fmt.Printf("%+v\n\n", todos)
	fmt.Println(time.Since(start).Seconds())

}

func GetUrl(Id int, source chan []byte) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/" + strconv.Itoa(Id))
	if err != nil {
		panic(err)
	}
	responseBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		panic(err)
	}
	source <- responseBody
}
