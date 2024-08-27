package main

import (
	"fmt"
	// "runtime/debug"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("the error is occured")
			// debug.PrintStack()
		fmt.Printf("error occured")
		}
	}()
	getUserInput()
	fmt.Println("end")

}

func getUserInput() {
	var value uint16
	fmt.Printf("enter number: ")
	fmt.Scanln(&value)
	if value == 1 {
		panic("number must not be 1")
	} else {
		fmt.Printf("your number is %d", value)
	}

}
