package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)
}
