package main

import (
	"fmt"
	"reflect" // to retrieve tags
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

type Human struct {
	Name   string `required max:"100"` // using tagging
	Origin string
}

func main() {
	aDoctor := Doctor{
		number:    4,
		actorName: "Jay Parmani",
		companions: []string{
			"Liz Shaw",
			"etc....",
		},
	}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName, aDoctor.companions[1])

	// Anonymous struct
	aDoctor2 := struct{ name string }{name: "Jay Parmani"}
	fmt.Println(aDoctor2)

	// Structs are by default passesd by value unlike maps and slices

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)

	c := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}

	fmt.Println(c)

	t := reflect.TypeOf(Human{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)

}
