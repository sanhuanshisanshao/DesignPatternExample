package main

import (
	"fmt"
	"strings"
	"sync"
)

type Command interface {
	execute() string
}

type Invoker struct {
	Command chan Command
}

func NewInvoker() *Invoker {
	i := &Invoker{Command: make(chan Command, 0)}
	return i
}

func (i *Invoker) Add(c Command) {
	i.Command <- c
}

func (i *Invoker) Call() {
	for {
		select {
		case c := <-i.Command:
			commands := c.execute()
			str := strings.Split(commands, "|")
			for _, v := range str {
				fmt.Printf("invoker receive command %v\n", v)
			}
			break
		default:
			//nothing to do
		}
	}
}

type Receiver struct {
	Commands []Command
}

func (r *Receiver) Add(c Command) {
	fmt.Printf("receiver receive a command\n")
	r.Commands = append(r.Commands, c)
}

func NewReceiver() *Receiver {
	return &Receiver{Commands: make([]Command, 0)}
}

type TVOpenCommand struct {
}

func NewTVOpenCommand(r *Receiver) *TVOpenCommand {
	t := &TVOpenCommand{}
	r.Add(t)
	return t
}

func (tc *TVOpenCommand) execute() string {
	fmt.Printf("command TV open command command1|command2|command3|command4|command5\n")
	return "command1|command2|command3|command4|command5"
}

//命令模式实现了对发送者和执行者（invoker）的解耦
//发送者只需要将命令发送给receiver，而不用管执行者如何执行命令
//执行者只需要从receiver接收命令，而不用管命令从何而来
func main() {
	wg := sync.WaitGroup{}

	receiver := NewReceiver()
	NewTVOpenCommand(receiver)
	invoker := NewInvoker()

	go func() {
		defer wg.Done()
		wg.Add(1)

		invoker.Call()
	}()

	for _, v := range receiver.Commands {
		invoker.Add(v)
	}

	wg.Wait()
}
