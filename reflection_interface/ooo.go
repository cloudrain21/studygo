package main

import "fmt"

type a struct {
	XX int
	YY int
}

type b struct {
	AA string
	XX int
}

type c struct {
	A a
	B b
}

func (A a) AAA() {
	fmt.Println("Function AAA() for A")
}

func (B b) AAA() {
	fmt.Println("Function AAA() for B")
}

func main() {
	var i c
	i.A.AAA()
	i.B.AAA()
}
