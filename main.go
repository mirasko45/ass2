package main

import "fmt"

type Component interface {
	elementPrice() float64
}

type ConcreteMethod struct {
}

func (epr *ConcreteMethod) elementPrice() float64 {
	return 64.0
}

type Decorator1 struct {
	second Component
}

func (pr *Decorator1) elementPrice() float64 {
	totalPrice := pr.second.elementPrice()
	return totalPrice + 64
}

type Decorator2 struct {
	third Component
}

func (pr *Decorator2) elementPrice() float64 {
	totalPrice := pr.third.elementPrice()
	return totalPrice + 60
}

func main() {

	action := &ConcreteMethod{}

	addSomething := &Decorator2{
		third: action,
	}

	addSomething2 := &Decorator1{
		second: addSomething,
	}

	fmt.Println(addSomething2.elementPrice())
}
