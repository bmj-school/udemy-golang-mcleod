package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

// Attach a method to the type person (increase method set!)
func (ptr_p *person) speak() {
	fmt.Println(ptr_p.first)
}

// A person *IS* a human (polymorphism)
type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{"Bill", "Bryson", 55}
	fmt.Println(p1)
	p1.speak()
	fmt.Printf("%v %T\n", p1, p1)
	saySomething(&p1)

	// p1 is a person:
	fmt.Println("Check if p1 is person:", interface{}(p1).(person))

	// p1 is NOT a human:
	fmt.Println("Check if p1 is human:", interface{}(p1).(human))
	if recover() != nil {
		fmt.Println("Recovering from panic!")
	}

	// But the ADDRESS is a human!
	fmt.Println("Check if p1 is human:", interface{}(&p1).(human))

}
