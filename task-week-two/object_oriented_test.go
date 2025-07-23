package main

import "testing"

func TestObjectOriented(t *testing.T) {
	rectangle := Rectangle{}
	circle := Circle{}
	rectangle.Perimeter()
	circle.Perimeter()
	circle.Area()
	rectangle.Area()
}

func TestEmployeePrintInfo(t *testing.T) {
	employee := Employee{
		Person: Person{
			Name: "张三",
			Age:  18,
		},
		EmployeeID: 1,
	}
	employee.PrintInfo()
}
