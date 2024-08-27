package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Person struct {
	Name   string `json:"first_name" xml:"name"`
	Age    uint   `json:"age,omitempty"`
	Family string `json:"family"`
	Gender string `json:"gender"`
}

func main() {
	person := Person{Name: "ali", Family: "alj", Gender: "male"}

	personJson, _ := json.MarshalIndent(person, "", "     ")
	personXml, _ := xml.MarshalIndent(person, "", "     ")

	fmt.Println(string(personJson))
	fmt.Println(string(personXml))
}
