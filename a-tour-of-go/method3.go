package main

import "fmt"

type employee struct {
	Name   string
	salary int
}

func (s *employee) AddSalary(param int) int {
	return s.salary + s.salary*param
}

func main() {
	em1 := new(employee)
	em1.Name = "john"
	em1.salary = 100

	fmt.Printf("Add double is :%d\n", em1.AddSalary(2))
}
