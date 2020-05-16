package main

import "fmt"

var x int = 32
var y string = "James"
var z bool = true

// type hotdog int
// var b hotdog

func main() {

	for i := 0; i < 10000; i++ {
		fmt.Printf("%v|", i)
	}

	for i := 65; i <= 90; i++ {
		fmt.Printf("%#U|", i)
	}

	bd := 1982
	for bd <= 2020 {
		fmt.Println(bd)
		bd++
	}

	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}

	x := "My String"
	if x == "asdf" {
		fmt.Println("NO")
	} else if x == "what" {
		fmt.Println("NO")
	} else {
		fmt.Println("Yep")
	}

	switch {
	case 1 == 1:
		fmt.Println("MATCH CASE")
	case 1 == 1:
		fmt.Println("MATCH CASE2")
	}
}
