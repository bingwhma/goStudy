package main

import "fmt"

type Human struct {
	name string
	age  int
}

type Student struct {
	Human
	School string
	loan   float32
}

type Employee struct {
	Human
	Company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("hi, i am %s. age:%d", h.name, h.age)
}

func (e Employee) SayHi() {
	fmt.Printf("hi, i am %s. age:%d", e.name, e.age)
}

type Men interface {
	SayHi()
}

func main() {
	mike := Student{Human{"mike", 25}, "zz", 0.00}
	sam := Employee{Human{"sam", 35}, "yyyy", 80.00}

	var i Men
	i = mike

	fmt.Println("\n mike is a student")
	i.SayHi()

	i = sam
	fmt.Println("\n sam is an Employee")
	sam.SayHi()
}
