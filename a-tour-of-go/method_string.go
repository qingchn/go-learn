package main

import (
	"fmt"
	"strconv"
)

/*
type TwoInts struct {
	a int
	b int
}
*/
type T struct {
	a int
	b float32
	c string
}

func (tn *T) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.FormatFloat(tn.b, 32) + "/" + strconv.Quote(tn.c) + ")"
}
func main() {
	/*
		two1 := new(TwoInts)
		two1.a = 12
		two1.b = 10
	*/
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("two1 is : %v\n", t)
	fmt.Println("two1 is:", t)
	fmt.Printf("two1 is: %T\n", t)
	fmt.Printf("two1 is: %#v\n", t)
}

/*
func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
*/
