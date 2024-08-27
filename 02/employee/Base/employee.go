package main

import "fmt"


const (
	FullTimeEmployeeTax = 0.1
	PartTimeEmployeeTax = 0.2
	ContractEmployeeTax = 0.3
)

type EmployeeCalculateSalary interface{
	CalculateSalary(hour int) (salary float64, tax float64)
}




type BaseEmployee struct{
	ID uint
	Name string
	NationalCode string
}


type FullTimeEmployee struct{
	BaseEmployee
	PayedEachDay float64

}

type ContractEmployee struct{
	BaseEmployee
	PayedBaseOnContract float64
}

type PartTimeEmployee struct{
	BaseEmployee
	PayedEachHour float64
}


func (employee FullTimeEmployee) CalculateSalary(hour int) (salary float64, tax float64){
	salary = float64(hour )* employee.PayedEachDay
	tax = FullTimeEmployeeTax * salary
	salary -= tax
	return 
}

func (employee PartTimeEmployee) CalculateSalary(hour int) (salary float64, tax float64){
	salary = float64(hour )* employee.PayedEachHour
	tax = FullTimeEmployeeTax * salary
	salary -= tax
	return 
}

func (employee ContractEmployee) CalculateSalary(hour int) (salary float64, tax float64){
	salary = float64(hour )* employee.PayedBaseOnContract
	tax = FullTimeEmployeeTax * salary
	salary -= tax
	return 
}

func main(){
	allEmployees := []EmployeeCalculateSalary{}

	PartTimeEmployees := []PartTimeEmployee{
		{BaseEmployee: BaseEmployee{ID:1, Name:"ali", NationalCode:"1223344"}, PayedEachHour:60000.0},
		{BaseEmployee: BaseEmployee{ID:1, Name:"ahmad", NationalCode:"2379424"}, PayedEachHour:40000.0},
	}

	FullTimeEmployees := []FullTimeEmployee{
		{BaseEmployee: BaseEmployee{ID:1, Name:"ali", NationalCode:"1223344"}, PayedEachDay:60000.0},
		{BaseEmployee: BaseEmployee{ID:1, Name:"rez", NationalCode:"2379424"}, PayedEachDay:40000.0},
	}

	ContractEmployees := []ContractEmployee{
		{BaseEmployee: BaseEmployee{ID:1, Name:"ali", NationalCode:"1223344"}, PayedBaseOnContract:60000.0},
		{BaseEmployee: BaseEmployee{ID:1, Name:"j", NationalCode:"2379424"}, PayedBaseOnContract:40000.0},
	}

	for _, employee := range PartTimeEmployees{
		allEmployees = append(allEmployees, employee)
	}

	for _, employee := range FullTimeEmployees{
		allEmployees = append(allEmployees, employee)
	}

	for _, employee := range ContractEmployees{
		allEmployees = append(allEmployees, employee)
	}

	for _, employee:= range allEmployees{
		salary, tax := employee.CalculateSalary(12)
		fmt.Printf("employee:(%T)\nsalary: %.0f, tax: %.0f\n\n", employee, salary, tax)
	}



}