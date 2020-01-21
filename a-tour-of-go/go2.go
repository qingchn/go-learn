package main

import "fmt"

type student struct {
	name string
}

func main() {
	var stuChan chan *student
	stuChan = make(chan *student, 10)

	stu := student{name: "stu001"}

	stuChan <- &stu
	var stu01 interface{}
	stu01 = <-stuChan

	var stu02 *student
	stu02, ok := stu01.(*student)
	if !ok {
		fmt.Println("can not convert")
		return
	}
	fmt.Println(stu02)
}
