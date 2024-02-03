package main

import "fmt"

type worker interface {
	getName() string
	getSalary() int
}

type employee struct {
	name   string
	salary int
}

type contractor struct {
	name        string
	hourlyRate  int
	hoursWorked int
}

func (e employee) getName() string {
	return e.name
}

func (e employee) getSalary() int {
	return e.salary
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyRate * c.hoursWorked
}

func test(worker worker) {
	fmt.Printf("Name is %s",worker.getName())
	fmt.Printf("Salary is %d",worker.getSalary())
}

func main() {
	newEmployee := employee{
		name: "Kevin",
		salary: 299,
	}

	newContractor := contractor{
		name:"Erin",
		hourlyRate: 10,
		hoursWorked: 20,
	}

	test(newEmployee)
	test(newContractor)
}