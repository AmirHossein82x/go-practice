package main

import "fmt"

type Person struct{
	Name string
	Age uint
	Family string
	Gender string
	
}


type PersonOptions func(*Person)

func main(){
	person1 := NewPerson(setName("ali"), setAge(22), setFamily("ghasemi"), setGender("male"))
	fmt.Printf("%+v", person1)
}


func NewPerson(options ...PersonOptions) *Person{
	person := &Person{}

	for _, option:= range options {
		option(person)
	}
	return person
}

func setName(name string) PersonOptions{
	return func(person *Person) {
		person.Name = name
	}
}

func setAge(age uint) PersonOptions {
	return func(person *Person) {
		person.Age = age
	}
}

func setFamily(family string) PersonOptions {
	return func(person *Person) {
		person.Family = family
	}
}

func setGender(gender string) PersonOptions {
	return func(person *Person) {
		person.Gender = gender
	}
}