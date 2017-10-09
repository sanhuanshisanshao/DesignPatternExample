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
}

type BasicDuck struct {
}

func (d *BasicDuck) Swim() {
	fmt.Println("basic duck swimming...")
}

func (d *BasicDuck) Eat() {
	fmt.Println("basic duck eatting...")
}

type RedheadDuck struct {
	//继承鸭子的通用接口
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
