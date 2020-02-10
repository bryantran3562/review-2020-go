package main

import (
	"fmt"
)

//BT - Interface defines a behavior of an object. It only specifies what methods an object should have and
//     it is up to the object to implement the details the methods.
// In go, interface is a set of methods signature, a type in which provided all the implemetion of those methods
// in the interface is said to be implement that interface.

type animal interface {
	speak() string
}

type dog struct {
}

func (d dog) speak() string {
	return "Woof"
}

type cat struct {
}

func (c cat) speak() string {
	return "Meow"
}

//BT - This function will accept any object that has method speak()
func animalSpeak(a animal) {
	fmt.Println(a.speak())
}

func main() {

	fmt.Println("Hi")

	d := dog{}

	c := cat{}

	fmt.Println(d.speak())
	fmt.Println(c.speak())

	fmt.Println("====================================")
	animalSpeak(d)
	animalSpeak(c)
}
