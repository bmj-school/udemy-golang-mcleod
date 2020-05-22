package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{"James", "Bonds", 32}
	p2 := person{"Bill", "Rogers", 56}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(string(bs))
	os.Stdout.Write(bs)
	myJson := `[{"First":"James","Last":"Bonds","Age":32},{"First":"Bill","Last":"Rogers","Age":56}]`

	s := []byte(myJson)
	fmt.Println(string(s))
	var loadedPeople []person

	err = json.Unmarshal(s, &loadedPeople)

	if err != nil {
		os.Exit(1)
	}

	fmt.Println("Reloaded the data:", loadedPeople)

}
