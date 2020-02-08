package main

import (
	"fmt"
)

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
