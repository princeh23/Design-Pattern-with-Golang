package main

import "fmt"

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

type Subject struct {
	observers []IObserver
}

func (sub *Subject) Register(observer IObserver) {
	sub.observers = append(sub.observers, observer)
}

func (sub *Subject) Remove(observer IObserver) {
	for i := 0; i < len(sub.observers); i++ {
		if sub.observers[i].GetName() == observer.GetName() {
			sub.observers = append(sub.observers[:i], sub.observers[i+1:]...)
			i--
		}
	}
}

func (sub *Subject) Notify(msg string) {
	for _, o := range sub.observers {
		o.Update(msg)
	}
}

type IObserver interface {
	Update(msg string)
	GetName() string
}

type Observer struct {
	Name string
}

func (Observer) Update(msg string) {
	fmt.Printf("Observer: %s\n", msg)
}

func (o Observer) GetName() string {
	return o.Name
}

func main() {
	obs1 := Observer{"a"}
	obs2 := Observer{"b"}
	obs3 := Observer{"c"}

	sub := &Subject{}
	sub.Register(obs1)
	sub.Register(obs2)
	sub.Register(obs3)
	sub.Notify("a b c all exists")
	sub.Remove(obs2)
	sub.Notify("a c exists")
}
