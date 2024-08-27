package main

import "fmt"

// anonymouse field
type Person struct {
	int
	string
	float32
}

func main(){
	// api call and get result

	apiResponse := struct {
		ResultCode int
		ResultMessage string
		TransactionAmount int
	}{
		ResultCode: 0,
		ResultMessage: "hello",
		TransactionAmount: 12,
	}

	fmt.Printf("%+v", apiResponse)

	person := Person{1, "alireza", 2.0}
	fmt.Printf("%+v", person)
}