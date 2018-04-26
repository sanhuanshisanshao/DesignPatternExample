package main

import "fmt"

//抽象主题角色
type Subject interface {
	request()
}

//代理主题角色
type Proxy struct {
	subject Subject
}

//真实主题角色
type RealSubject struct {
}

func NewProxy() *Proxy {
	return &Proxy{
		subject: &RealSubject{},
	}
}

func (p *Proxy) request() {
	p.preRequest()
	p.subject.request()
	p.afterRequest()
}

func (p *Proxy) afterRequest() {
	fmt.Printf("proxy subject after request\n")
}

func (p *Proxy) preRequest() {
	fmt.Printf("proxy subject pre request\n")
}

func (rp *RealSubject) request() {
	fmt.Printf("real subject do request\n")
}

func main() {

	proxy := NewProxy()
	proxy.request()

}
