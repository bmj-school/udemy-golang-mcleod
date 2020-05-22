package main

import "fmt"

type person struct {
	age int
}

func changeMe(p *person) {
	fmt.Println(p)
	(*p).age = 10
}

func main() {
	x := 42
	fmt.Println(&x)

	p1 := person{32}
	fmt.Println(p1)

	changeMe(&p1)

	fmt.Println(p1)
}
