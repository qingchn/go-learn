package main

import (
	"fmt"
)

/*
type Engine struct {
	Start()
    Stop()
}
*/

type Car struct {
	wheelCount int
	//	Engine
}

func (c Car) NumberOfWheels() int {
	return c.wheelCount
}

type Mercedes struct {
	Car
	Name string
}

func (m Mercedes) sayHiToMerkel() string {
	return m.Name
}

func main() {
	m1 := Mercedes{Car{4}, "Mercedes"}
	fmt.Println(m1.NumberOfWheels())
	fmt.Println(m1.sayHiToMerkel())
}
