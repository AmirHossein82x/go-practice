package main

import "fmt"


type Person struct {
	Age int
}

func main(){
	my := []int{1,2, 3,4,5}
	persons := []Person{Person{Age: 20}, Person{Age: 18}}
	change(my)
	change2(persons)
	fmt.Println(my)
	fmt.Printf("%+v", persons)
}

func change(my []int){
	for i, item:= range my{
		my[i] = item * 3
	}
}

func change2(persons []Person){
	persons[0] = Person{Age: 33}
}