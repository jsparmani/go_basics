package main

import (
	"fmt"
)

func main() {
	var n bool = true
	fmt.Printf("%v, %T\n", n, n)

	n = 1 == 2
	fmt.Printf("%v, %T\n", n, n)

	// By default a variable is initialised to 0
	var y int
	fmt.Printf("%v, %T\n", y, y)

	// Data-types
	// int, bool, uint8, uint16, uint32, float32, float64

	// complex64, complex128

	var j complex64 = 1 + 2i
	var m complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", m, m)

	fmt.Printf("%v, %T\n", real(j), real(j))
	fmt.Printf("%v, %T\n", imag(j), imag(j))

	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// var c int8 = 3
	// fmt.Println(a + c) -- Not allowed because a and c have different datatypes, use explicit type conversion

	// text types
	s := "this is a string" // utf8 character
	fmt.Printf("%v, %T\n", s[2], s[2])

	// strings can be concatenated

	// converting string to an array of bytes
	t := []byte(s)
	fmt.Printf("%v, %T\n", t, t)

	r := 'a' // alias for int32
	fmt.Printf("%v, %T\n", r, r)

}
