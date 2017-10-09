package main

import "fmt"

type Animal interface {
	makeSound()
}

type Dog struct {
}

func (d *Dog) makeSound() {
	fmt.Println("dog say wang wang ...")
}

func NewDog() *Dog {
	return &Dog{}
}

type Cat struct {
}

func NewCat() *Cat {
	return &Cat{}
}

func (c *Cat) makeSound() {
	fmt.Println("cat say miao miao miao ...")
}

func main() {

	var a Animal
	a = NewDog()
	a.makeSound()

	a = NewCat()
	a.makeSound()

}
