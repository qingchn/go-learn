package main

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
