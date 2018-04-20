package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

type Colleague interface {
	ReceiveMsg()
	SendMsg(who, msg string)
	SetMediator(m Mediator)
	Push(msg string)
}

type Mediator interface {
	Operation(who string, s string) error
	Register(who string, col Colleague) error
}

type ConcreteMediator struct {
	Colleagues map[string]Colleague
}

func NewMediator() Mediator {
	return &ConcreteMediator{make(map[string]Colleague)}
}

func (cm *ConcreteMediator) Operation(who, s string) error {
	val, ok := cm.Colleagues[who]
	if !ok {
		return errors.New("can't find colleague from mediator")
	}
	go val.Push(s)
	return nil
}

func (cm *ConcreteMediator) Register(who string, col Colleague) error {
	_, ok := cm.Colleagues[who]
	if ok {
		return errors.New("colleague exist in mediator")
	}
	cm.Colleagues[who] = col
	return nil
}

type ConcreteColleague struct {
	Msgs     chan string
	Mediator Mediator
}

func NewColleague() Colleague {
	return &ConcreteColleague{
		Msgs: make(chan string, 100)}
}

func (cc *ConcreteColleague) SendMsg(who, msg string) {
	cc.Mediator.Operation(who, msg)
}

func (cc *ConcreteColleague) ReceiveMsg() {
	for true {
		select {
		case m := <-cc.Msgs:
			fmt.Printf("receive msg %s", m)
		}
	}
}

func (cc *ConcreteColleague) Push(msg string) {
	cc.Msgs <- msg
}

func (cc *ConcreteColleague) SetMediator(m Mediator) {
	cc.Mediator = m
}

func main() {

	wg := sync.WaitGroup{}

	mediator := NewMediator()

	coleagueA := NewColleague()
	coleagueB := NewColleague()

	mediator.Register("coleagueA", coleagueA)
	mediator.Register("coleagueB", coleagueB)

	coleagueA.SetMediator(mediator)
	coleagueB.SetMediator(mediator)

	go func() {
		wg.Add(1)
		defer wg.Done()
		for {
			coleagueA.SendMsg("coleagueB", "hello world!\n")
			time.Sleep(1 * time.Second)
		}
	}()

	go coleagueB.ReceiveMsg()

	go coleagueA.ReceiveMsg()

	wg.Wait()
}
