package main

type Simpler interface {
	Get() int
	Set() int
}

type Simple struct {
	a int
	b int
}

func (s Simple) Get() int {
	return s.a
}

func (s Simple) Set(s int) int {
	return s.a + s
}

func 