package main

import (
	"errors"
	"fmt"
	"strings"
)

type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

var employees []Employee

func addEmployee(id int, name string, age int, department string) error {
	
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("employee ID must be unique")
		}
	}

	
	if age <= 18 {
		return errors.New("employee age must be greater than 18")
	}

	employees = append(employees, Employee{ID: id, Name: name, Age: age, Department: department})
	return nil
}

func searchEmployeeByID(id int) (*Employee, error) {
	for _, emp := range employees {
		if emp.ID == id {
			return &emp, nil
		}
	}
	return nil, errors.New("employee not found")
}

func searchEmployeeByName(name string) (*Employee, error) {
	for _, emp := range employees {
		if strings.EqualFold(emp.Name, name) {
			return &emp, nil
		}
	}
	return nil, errors.New("employee not found")
}

func listEmployeesByDepartment(department string) []Employee {
	var result []Employee
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			result = append(result, emp)
		}
	}
	return result
}

func countEmployees(department string) int {
	count := 0
	for _, emp := range employees {
		if strings.EqualFold(emp.Department, department) {
			count++
		}
	}
	return count
}

func main() {
	const (
		HR = "HR"
		IT = "IT"
		SALES = "Sales"
	)

	
	if err := addEmployee(1, "Alice", 25, HR); err != nil {
		fmt.Println("Error:", err)
	}
	if err := addEmployee(2, "Bob", 30, IT); err != nil {
		fmt.Println("Error:", err)
	}
	if err := addEmployee(3, "Charlie", 22, SALES); err != nil {
		fmt.Println("Error:", err)
	}
	if err := addEmployee(2, "David", 28, HR); err != nil { 
		fmt.Println("Error:", err)
	}

	
	if emp, err := searchEmployeeByID(1); err == nil {
		fmt.Println("Found employee by ID:", *emp)
	} else {
		fmt.Println("Error:", err)
	}

	if emp, err := searchEmployeeByName("Bob"); err == nil {
		fmt.Println("Found employee by name:", *emp)
	} else {
		fmt.Println("Error:", err)
	}

	
	fmt.Println("Employees in HR department:", listEmployeesByDepartment(HR))

	
	fmt.Printf("Number of employees in IT department: %d\n", countEmployees(IT))
}
