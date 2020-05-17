package main

import "fmt"

func main() {
	myArray := [...]float32{3, 3, 2.0, 1.1}
	fmt.Println(myArray)
	for i, v := range myArray {
		fmt.Printf("%v %v\n", i, v)
	}

	x := []int{2, 2, 3, 4, 99, 200}
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	mySlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(mySlice[1:6])

	mySlice2 := append(mySlice, 53)
	fmt.Println(mySlice2)

	mySlice3 := append(mySlice, mySlice...)
	fmt.Println(mySlice3)

	// Delete the 3rd value

	mySlice4 := append(mySlice[0:2], mySlice[3:]...)
	fmt.Println("Del", mySlice4)

	delTarget := 3
	mySlice4 = append(mySlice[0:delTarget], mySlice[delTarget+1:]...)
	fmt.Println("Del", mySlice4)

	states := make([]string, 50)
	states = []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado",
		"Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana",
		"Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan",
		"Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire",
		"New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma",
		"Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee",
		"Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}

	for i, v := range states {
		fmt.Println(i, v)
	}

	pops := []int{
		4903185, 731545, 7278717, 3017825, 39512223, 5758736, 3565287, 973764, 21477737,
		10617423, 1415872, 1787065, 12671821, 6732219, 3155070, 2913314, 4467673, 4648794, 1344212,
		6045680, 6949503, 9986857, 5639632, 2976149, 6137428, 1068778, 1934408, 3080156, 1359711, 8882190,
		2096829, 19453561, 10488084, 762062, 11689100, 3956971, 4217737, 12801989, 1059361, 5148714,
		884659, 6833174, 28995881, 3205958, 623989, 8535519, 7614893, 1792147, 5822434, 578759}

	for i, v := range pops {
		fmt.Println(i, v)
	}

	fmt.Println("Excercise")
	statesMap := make(map[string]int, 50)

	for i := 0; i < 50; i++ {
		// fmt.Println(states[i], pops[i])
		statesMap[states[i]] = pops[i]
	}

	for k, v := range statesMap {
		fmt.Println(k, v)
	}

	if _, ok := statesMap["asdf"]; !ok {
		fmt.Println("This key does not exist!!!")
	}

	delete(statesMap, "Arizona")

	if _, ok := statesMap["Arizona"]; !ok {
		fmt.Println("This key does not exist!!!")
	}


	// for k, v := range statesMap{
	// 	// statesMap[states
	// 	fmt.Println(k,v)
	// }

	// table := [][]string{states, pops}
}
