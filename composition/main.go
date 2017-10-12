package main

import "fmt"

type Fly interface {
	Fly()
}

type slowFly struct {
}

func (f *slowFly) Fly() {
	fmt.Println("fly slowly ...")
}

type fastFly struct {
}

func (f *fastFly) Fly() {
	fmt.Println("fly quickly ...")
}

type Quack interface {
	Quack()
}

type gentlyQuack struct {
}

func (q *gentlyQuack) Quack() {
	fmt.Println("quack gently ...")
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
	Q Quack
	F Fly
	D Duck
}

func (d *RedheadDuck) setFly(style string) {
	switch style {
	case "slow":
		d.F = &slowFly{}
		break
	default:
		d.F = &fastFly{}
		break
	}
}

func (d *RedheadDuck) setQuack() {
	d.Q = &gentlyQuack{}
}

func main() {
	duck := &RedheadDuck{}
	duck.D = &BasicDuck{}
	duck.setFly("slow")
	duck.setQuack()

	duck.D.Eat()
	duck.D.Swim()
	duck.F.Fly()
	duck.Q.Quack()

	duck.setFly("")

	duck.D.Eat()
	duck.D.Swim()
	duck.F.Fly()
	duck.Q.Quack()

}
