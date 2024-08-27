package main

import "fmt"


type Person struct{
	Name string
	Age uint
	Family string
	Gender string
	
}


type PersonBuilder struct {
	Person
}

func main(){
	builder := &PersonBuilder{}
	person1:= builder.setName("ali").setAge(12).setFamily("hosseini").setGender("male").create()
	person2:= builder.setName("ahmamd").setAge(15).setFamily("hosseini").setGender("male").create()
	fmt.Printf("%+v\n", person1)
	fmt.Printf("%+v", person2)
	

}


func (builder *PersonBuilder) setName(name string) *PersonBuilder{
	builder.Name = name
	return builder
}


func (builder *PersonBuilder) setFamily(family string) *PersonBuilder{
	builder.Family = family
	return builder
}

func (builder *PersonBuilder) setAge(age uint) *PersonBuilder{
	builder.Age = age
	return builder
}

func (builder *PersonBuilder) setGender(gender string) *PersonBuilder{
	builder.Gender = gender
	return builder
}

func (builder *PersonBuilder) create() Person{
	return builder.Person
}


