package main

import "fmt"

var myInt int = 32
var myFloat float64 = 32.3
const (
	a = 542
	b int = 322
)

func myprint(astring string) {
	astring += "\n"
	fmt.Printf(astring)
}


const (
	c1 = iota
	c2
	c3
	c4
)
func main() {
	fmt.Println("Excercise 1")
	fmt.Printf("int %d, float %f, hex %#x, bin %#b\n",myInt, myInt, myInt, myInt)
	fmt.Printf("int %d, float %.2f, hex %#x, bin %#b\n", myFloat, myFloat, myFloat, myFloat)

	fmt.Println("Excercise 2 - skip")

	fmt.Println("Excercise 3 - skip")
	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", b, b)

	fmt.Println("Excercise 4")
	a := 42
	fmt.Printf("int %d, float %f, hex %#x, bin %#b\n", a, a, a, a)
	ashift := a << 1
	fmt.Printf("int %d, float %f, hex %#x, bin %#b\n", ashift, ashift, ashift, ashift)

	fmt.Println("Excercise 5")
	myLiteral := `String
	Lit		x\n`
	fmt.Printf(myLiteral)

	fmt.Printf("\n%v %v %v", c1,c2,c3)



}