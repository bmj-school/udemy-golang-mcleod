package main

import "fmt"

func main() {
	s := "String here"
	fmt.Println(s)

	for i:=0; i<len(s); i++ {
		fmt.Printf("%#U\n", s[i])
	}

	// For each style
	for _, e:= range s{
		fmt.Printf("%#U\n", e)
	}
}