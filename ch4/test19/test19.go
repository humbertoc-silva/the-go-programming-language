package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func EmployeeByID(id int) *Employee {
	var employee Employee
	employee.ID = 2
	employee.Name = "Humberto"
	employee.Address = "Fifth Avenue, 1000"
	employee.DoB = time.Now()
	employee.Position = "Software Engineer"
	employee.Salary = 1000
	employee.ManagerID = 1

	return &employee
}

func main() {
	var dilbert Employee
	dilbert.Salary -= 5000 // demoted, for writing too few lines of code

	// access a field through a pointer
	position := &dilbert.Position
	*position = "Senior " + *position // promoted, for outsourcing to Elbonia
	fmt.Println(dilbert)

	// dot notation with pointer
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)" // this is equivalent to: (*employeeOfTheMonth).Position += " (proactive team player)"
	fmt.Println(dilbert)

	employee := EmployeeByID(2)
	fmt.Println(employee.Position)
	employee.Salary = 0 // fired for... no real reason
	fmt.Println(employee)
}
