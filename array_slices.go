package main

import (
	"fmt"
)

func main() {
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

	// Slice in Go

	d := []int{1, 2, 3}
	fmt.Printf("Length: %v\n", len(d))
	fmt.Printf("Capacity: %v\n", cap(d))
}
