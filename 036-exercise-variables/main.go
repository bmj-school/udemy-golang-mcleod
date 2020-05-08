package main

import "fmt"

var x int = 32
var y string = "James"
var z bool = true
// type hotdog int
// var b hotdog

func main() {
	s := fmt.Sprintf("%v %v %v",x,y,z)
	fmt.Println(s)
}