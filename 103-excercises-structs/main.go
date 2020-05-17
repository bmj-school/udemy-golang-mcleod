package main

import "fmt"

type person struct {
	first       string
	last        string
	favFlavours []string
}

func main() {
	p1 := person{
		first:       "Joe",
		last:        "McSnugglet",
		favFlavours: []string{"vanilla", "berry"},
	}
	p2 := person{
		first:       "Sue",
		last:        "McSnugglet",
		favFlavours: []string{"vanilla", "berry"},
	}

	fmt.Println(p1)

	m := map[string]person{}
	m[p1.last+p1.first] = p1
	m[p2.last+p2.first] = p2

	for k, v := range m {
		fmt.Println(k, v)
	}

	// Embedded structs
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	t1 := truck{
		vehicle:   vehicle{doors: 2, color: "blue"},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{doors: 4, color: "white"},
		luxury:  false,
	}
	fmt.Println(t1)
	fmt.Println(s1)

	// Anonymous struct
	s:= struct {
		first string
		friends map[string]int
	}{
		first: "James",
		friends: map[string]int {
			"Bill": 333,
			"sue": 3311111,
		},
	}
	fmt.Println(s)
}
