package main

import (
	"fmt"
)

const a int16 = 27

const (
	f = iota
	g = iota
	h = iota
)

const (
	i = iota
)

func main() {
	const myConst int = 42
	// myConst = 27 cannot be done
	fmt.Printf("%v, %T\n", myConst, myConst)

	// Shadowing
	const a int = 14
	fmt.Printf("%v, %T\n", a, a)

	var b int = 27
	fmt.Println(a + b)

	const d = 100
	fmt.Printf("%v, %T\n", d, d)

	// Enumerated constants
	const e = iota
	fmt.Printf("%v, %T\n", e, e)

	// iota
	fmt.Println(f, g, h)
	fmt.Println(i)

}
