package main

import "fmt"

type Strategy interface {
	algorithm(a, b int)
}

type Context struct {
	S Strategy
}

func (c *Context) algorithm(a, b int) {
	c.S.algorithm(a, b)
}

func (c *Context) setAlgorithm(s Strategy) {
	c.S = s
}

type ConcreteStrategyA struct {
}

func (cs *ConcreteStrategyA) algorithm(a, b int) {
	fmt.Printf("%v + %v = %v\n", a, b, (a + b))
}

type ConcreteStrategyB struct {
}

func (cs *ConcreteStrategyB) algorithm(a, b int) {
	fmt.Printf("%v * %v = %v\n", a, b, (a * b))
}

func main() {
	context := &Context{}
	concreteA := &ConcreteStrategyA{}
	concreteB := &ConcreteStrategyB{}

	context.setAlgorithm(concreteA)
	context.algorithm(3, 4)

	context.setAlgorithm(concreteB)
	context.algorithm(3, 4)
}
