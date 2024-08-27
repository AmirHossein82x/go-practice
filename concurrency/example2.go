package main

import (
	"fmt"
	"math/rand"
)

type Employee struct {
	Id int
	Salary int64
}


func main(){
	money := 150000000000
	myEmployee := []Employee{}
	for i:= 1; i < 1000; i ++ {
		myEmployee = append(myEmployee, Employee{Id: i, Salary: getRandomSalary()})
	}
	// fmt.Printf("%+v", myEmployee)
	
	PaySalary(&money, &myEmployee)
	// fmt.Printf("%+v", myEmployee)
	fmt.Println(money)
	fmt.Println(myEmployee)
	
}

func getRandomSalary()int64 {
	return int64(rand.Intn(10000))
}


func PaySalary(money *int, employees *[]Employee){
	for idx, employee:= range *employees{
		if employee.Salary < int64(*money) {
			*money -= int(employee.Salary)
			employees[idx] = Employee{Id: employee.Id, Salary: employee.Salary}
		}else{
			fmt.Println("not enough money")
		}

	}
}