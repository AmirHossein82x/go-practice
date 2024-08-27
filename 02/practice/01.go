package main

import (
	"fmt"
	"unicode"
)

func main() {
	var names []string = []string{}
	var name string
	for i := 0; i < 3; i++ {
		fmt.Scanln(&name)
		names = append(names, name)
	}
	fmt.Println(names)
	result := FilterLowers(names)
	fmt.Println(result)

}


func FilterLowers(names []string) []string{
	newSlc := []string{}
	for _, name:= range names{
		if IsLower(name){
			newSlc = append(newSlc, name)
		}
	}
	return newSlc
}



func IsLower(name string) bool {
	for _, chr:= range name{
		if !unicode.IsLower(chr) && unicode.IsLetter(chr){
				return false
		}
	}
	return true
}