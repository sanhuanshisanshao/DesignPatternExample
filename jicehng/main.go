package main

import "fmt"

type Fly interface {
	Fly()
}

type Quack interface {
	Quack()
}

type Duck interface {
	Swim()
	Eat()
	//Name(name string)
}

type BasicDuck struct {
}

func (d *BasicDuck) Swim() {
	fmt.Println("basic duck swimming...")
}

func (d *BasicDuck) Eat() {
	fmt.Println("basic duck eatting...")
}

//func (d *BasicDuck) Name(name string) {
//	fmt.Println("duck name is " + name)
//}

type RedheadDuck struct {
	D Duck
}

func (d *RedheadDuck) Fly() {
	fmt.Println("red head duck flying...")
}

func main() {
	var d RedheadDuck
	d.D = &BasicDuck{}

	d.D.Eat()
	d.D.Swim()
	d.Fly()
}
