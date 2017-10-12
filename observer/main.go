package main

import "container/list"

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
