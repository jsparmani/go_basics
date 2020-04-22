package main

import (
	"fmt"
)

func main() {

	statePopulations := make(map[string]int)

	statePopulations = map[string]int{
		"California": 5634637,
		"Texas":      2135980,
		"Florida":    144749,
		"New York":   25236,
		"Illinois":   68973,
		"Ohio":       1574758,
	}

	// Slices cannot be used as keys because they aren't equated

	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Ohio"])

	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)

	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

	fmt.Println(statePopulations["Georgia"])
	fmt.Println(statePopulations)

	pop, ok := statePopulations["Oho"]
	// ok returns false if key is not found
	fmt.Println(pop, ok)
	fmt.Println(len(statePopulations))

	// passing maps is like pointers
}
