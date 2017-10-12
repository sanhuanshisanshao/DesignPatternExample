package main

import (
	"container/list"
	"fmt"
)

type Observer interface {
	Update(s string)
}

type Subject interface {
	registerObserver(o Observer)
	removeObserver(o Observer)
	notifyObserver()
}

type ScoreSubject struct {
	Result       string
	ObserverList *list.List
}

func (s *ScoreSubject) registerObserver(o Observer) {
	s.ObserverList.PushBack(o)
}

func (s *ScoreSubject) removeObserver(o Observer) {
	for e := s.ObserverList.Front(); e != nil; e = e.Next() {
		if e.Value.(Observer) == o {
			s.ObserverList.Remove(e)
		}
	}
}

func (s *ScoreSubject) notifyObserver() {
	for e := s.ObserverList.Front(); e != nil; e = e.Next() {
		t := e.Value.(Observer)
		t.Update(s.Result)
	}
}

func (s *ScoreSubject) SetScoreResult(score string) {
	s.Result = score
	s.notifyObserver()
}

type CurrentObserver struct {
	ScoreResult string
}

func NewCurrentObserver(s *ScoreSubject) *CurrentObserver {
	o := &CurrentObserver{ScoreResult: "0-0"}
	s.registerObserver(o)
	return o
}

func (o *CurrentObserver) Update(s string) {
	o.ScoreResult = s
}

func (o *CurrentObserver) Display() {
	fmt.Printf("current observer %p score result is %v\n", o, o.ScoreResult)
}

func main() {
	scoreSubject := &ScoreSubject{Result: "0-0", ObserverList: list.New()}

	//为主题注册观察者
	observer1 := NewCurrentObserver(scoreSubject)
	observer2 := NewCurrentObserver(scoreSubject)

	//打印观察者的当前值
	observer1.Display()
	observer2.Display()

	//主题更新数据
	scoreSubject.SetScoreResult("1-0")
	observer1.Display()
	observer2.Display()

	scoreSubject.SetScoreResult("1-1")
	observer1.Display()
	observer2.Display()

	//observer1 取消订阅主题
	scoreSubject.removeObserver(observer1)
	scoreSubject.SetScoreResult("2-1")
	//observer1不再获取主题的最新推送
	observer1.Display()
	observer2.Display()
}
