package main

import "fmt"

type Employee interface {
	findSalary() int
}

type FullTime struct {
	daysWorked int
}

type Contractor struct {
	daysWorked int
}
type Freelancer struct {
	 hrsWorked int
}
func (f FullTime) findSalary() int {
	return f.daysWorked*500
}
func (c Contractor) findSalary() int {
	return c.daysWorked*100
}
func (fl Freelancer) findSalary() int {
	return fl.hrsWorked*10
}

func main() {
	fte:= FullTime{25}
	cont:=Contractor{18}
	freeLan:=Freelancer{44}
	emp := []Employee{fte,cont,freeLan}
	for _,i := range emp{
		fmt.Println(i.findSalary())
	}
}
