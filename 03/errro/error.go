package main

import (
	"errors"
	"fmt"
)

func main() {
	err, result := getUserInput()
	if err != nil{
		fmt.Println("the error occured")
	}else {
		fmt.Println(result)
	}
	fmt.Println("end")

}

func getUserInput() (err error, result string) {
	var value int
	fmt.Printf("enter your number: ")
	fmt.Scanln(&value)
	if value == 1 {
		return errors.New("you should not enter 1"), "1"
	} else {
		return nil, "your number is " + string(result)
	}

}
