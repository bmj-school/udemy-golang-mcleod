package main

import "fmt"

type SQUARE struct {
	L int
	W int
}

type CIRCLE struct {
	R float32
}


func (s SQUARE) AREA() float32 {
	return float32(s.L * s.W)
}

func (c CIRCLE) AREA() float32 {
	return float32(2.0) * float32(3.14) * c.R
}

type SHAPE interface {
	AREA() float32
}

func INFO(s SHAPE) float32 {
	// fmt.Println(s)
	return s.AREA()
}

func main() {
	s := SQUARE{2, 2}
	fmt.Println(INFO(s))

	c := CIRCLE{3}
	fmt.Println(INFO(c))
}
