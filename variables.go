package main

import (
	"fmt"
	"strconv"
)

//Variables can also be declared here
var m int = 12

// Block declaration

var (
	actorName string = "Jay"
	season    int    = 11
)

var (
	counter int = 0
	i       int = 0
)

// J exported
var J int = 23

func main() {
	var i int = 42
	i = 42
	var j float32 = 27
	k := 99
	fmt.Println(i, j, k)
	fmt.Printf("%v, %T\n", j, j)

	// Converting one variable type to other
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)

	// Converting a number to it's unicode character
	var y int = 42
	var z string
	z = string(y)
	fmt.Println(z)

	// To convert as it is in same form
	z = strconv.Itoa(y)
	fmt.Println(z)
}
