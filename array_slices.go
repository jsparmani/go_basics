package main

import (
	"fmt"
)

func main() {

	// Arrays don't work as pointers by default
	grades := [3]int{97, 85, 93}
	grades2 := [...]int{97, 85, 93, 10}
	fmt.Printf("Grades : %v, %T\n", grades, grades)
	fmt.Printf("Grades : %v, %T\n", grades2, grades2)

	var students [3]string
	fmt.Printf("Students, %v\n", students)
	students[0] = "Lisa"
	students[2] = "Ahmed"
	students[1] = "Arnold"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student 1: %v\n", students[1])
	fmt.Printf("Number of students: %v\n", len(students))

	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	// Slice in Go  - It's like dynamic arrays
	// Slices work like pointers by default

	z := []int{1, 2, 3}
	fmt.Printf("Length: %v\n", len(z))
	fmt.Printf("Capacity: %v\n", cap(z))

	// More Slice usage
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   // slice of all elements
	c := a[3:]  // slice from 4th element to end
	d := a[:6]  // slice first 6 elements
	e := a[3:6] // slice 4th, 5th and 6th elements
	a[5] = 42
	b[5] = 21
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// Using the built in make function
	f := make([]int, 3, 100)
	fmt.Println(f)
	fmt.Printf("Length: %v\n", len(f))
	fmt.Printf("Capacity: %v\n", cap(f))

	f = append(f, 1)
	fmt.Println(f)
	fmt.Printf("Length: %v\n", len(f))
	fmt.Printf("Capacity: %v\n", cap(f))

}
