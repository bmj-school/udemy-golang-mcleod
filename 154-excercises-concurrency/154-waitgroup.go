package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		fmt.Println("T1")
		wg.Done()
	}()
	go func() {
		fmt.Println("T2")
		wg.Done()
	}()
	go func() {
		fmt.Println("T3")
		wg.Done()
	}()
	wg.Wait()
}
