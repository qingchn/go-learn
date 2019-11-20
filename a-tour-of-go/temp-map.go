package main

import "fmt"

/*
func main() {
	sm := make([]map[int]string, 5)
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)
}
*/
/*
func main() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)
}
*/
/*
func main() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	fmt.Println(m1)
	m2 := make(map[string]int)

	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)
}

*/

/*

func main() {
	A(1, 2, 3, 4, 5, 6, 7, 8)
}

func A(a ...int) {
	fmt.Print(a)
}

*/

// 闭包
/*
func main() {
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}


func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
*/
/*
func main() {
	fmt.Printf("a")
	defer fmt.Println("b \n")
	defer fmt.Println("c")
}

*/
/*
func main() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}
*/
/*
func main() {
	for i :=0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}
*/
/*
func main() {
	A()
	B()
	C()

}

func A() {
	fmt.Println("Func A")
}

func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in B")
}

func C() {
	fmt.Println("Func C")
}

*/
/*
func main() {
	var min, max int
	min, max = MinMax(78, 65)
	fmt.Printf("Minmium is : %d, Maximum is: %d\n", min, max)

}

func MinMax(a int, b int) (min int, max int) {
	if a < b {
		min = a
		max = b
	} else {
		min = b
		max = a
	}
	return a, b
}
*/
/*
func main() {
	x := min(1, 3, 2, 0)
	fmt.Printf("The minimum is : %d \n", x)
	slice := []int{7, 9, 3, 5, 1}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
*/
/*
func main() {
	doDBOperations()
}

func connectToDB() {
	fmt.Println("OK, connected to db")
}

func disconnectFromDB() {
	fmt.Println("OK, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB() //function called here with defer
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return //terminate the program
}
*/

/*

func trace(s string) {
	fmt.Println("entering:", s)
}

func untrace(s string) {
	fmt.Println("leaving: ", s)
}

func a() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	a()
}

func main() {
	b()
}
*/
/*
func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

func main() {
	func1("Go")
}
*/
/*
func main() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is : %d \n", i, result)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
*/

/*
func main() {
	f()
}

func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d", i) }
		g(i)
		fmt.Printf("- g if of type %T and has value %v\n", g, g)
	}
}
*/
/*
func main() {
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v \n", TwoAdder(3))

}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
*/
/*
type person struct {
	Name string
	Age  int
}

func main() {
	a := person{}
	fmt.Println(a)
}
*/


type Shaper interface {
	Area () float32
}


type Square struct {
	side float32
}


func (sq *Square) Area() float32 {
	return sq.side * sq.side
}


func main() {
	sq1 := new(Square)
	sq1.side = 5

    var areaIntf Shaper
	areaIntf = sq1
	// shorter,without separate declaration:
    // areaIntf := Shaper(sq1)
    // or even:
	// areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())

}