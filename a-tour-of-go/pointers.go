package main

import "fmt"

/*
func main() {
	i, j := 42, 2701
	p := &i        //point to i
	fmt.Println(p) // read i through the pointer
	fmt.Println(&i)
	fmt.Println(&p)
	*p = 21        // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j
	*p = *p / 37

	fmt.Println(j)

}
*/

/*
func main() {
	var house = "Malibu Point 10880, 90265"
	prt := &house
	fmt.Printf("ptr type: %T\n", prt)

	fmt.Printf("address: %p\n", prt)

	value := *prt
	fmt.Printf("value type: %T\n", value)
	fmt.Printf("varlue: %s\n", value)
}
*/

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)
}
