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


type EmployeeBuilder struct {
	BaseEmployee
}

type PartTimeBuilder struct {
	PartTimeEmployee
}

type FullTimeBuilder struct {
	FullTimeEmployee
}

type ContractTimeBuilder struct {
	ContractEmployee
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
	fmt.Println("full time", tax, FullTimeEmployeeTax)
	return 
}

func (employee PartTimeEmployee) CalculateSalary(hour int) (salary float64, tax float64){
	salary = float64(hour )* employee.PayedEachHour
	tax = FullTimeEmployeeTax * salary
	salary -= tax
	fmt.Println("part time", tax, PartTimeEmployeeTax)
	return 
}

func (employee ContractEmployee) CalculateSalary(hour int) (salary float64, tax float64){
	salary = float64(hour )* employee.PayedBaseOnContract
	tax = FullTimeEmployeeTax * salary
	salary -= tax
	fmt.Println("contract time", tax, ContractEmployeeTax)
	return 
}

func main(){
	allEmployees := []EmployeeCalculateSalary{}
	builder := &EmployeeBuilder{}

	allEmployees = append(allEmployees, builder.setId(1).setName("ali").setNationalCode("1234566").fullTime(200000))
	allEmployees = append(allEmployees, builder.setId(2).setName("teza").setNationalCode("55555").partTime(200000))
	allEmployees = append(allEmployees, builder.setId(3).setName("sflsj").setNationalCode("777777").contractTime(200000))
	
	for _, em := range allEmployees{
		salary, tax := em.CalculateSalary(12)
		fmt.Printf("employee(%T), salary: %f,   tax: %f\n", em, salary, tax)
	}



}




func (builder *EmployeeBuilder) setId(id uint) *EmployeeBuilder{
	builder.ID = id
	return builder
}

func (builder *EmployeeBuilder) setName(name string) *EmployeeBuilder{
	builder.Name = name
	return builder
}

func (builder *EmployeeBuilder) setNationalCode(nationalCode string) *EmployeeBuilder{
	builder.NationalCode = nationalCode
	return builder
}


func (builder *EmployeeBuilder) fullTime(payedEachDay uint) FullTimeEmployee{
	var fullTimeEmployee FullTimeEmployee = FullTimeEmployee{BaseEmployee: builder.BaseEmployee, PayedEachDay: float64(payedEachDay)}
	return fullTimeEmployee
}

func (builder *EmployeeBuilder) partTime(payedEachHour uint) PartTimeEmployee{
	var partTimeEmployee PartTimeEmployee = PartTimeEmployee{BaseEmployee: builder.BaseEmployee, PayedEachHour: float64(payedEachHour)}
	return partTimeEmployee
}

func (builder *EmployeeBuilder) contractTime(payedBaseOnContract uint) ContractEmployee{
	var fullTimeEmployee ContractEmployee = ContractEmployee{BaseEmployee: builder.BaseEmployee, PayedBaseOnContract: float64(payedBaseOnContract)}
	return fullTimeEmployee
}