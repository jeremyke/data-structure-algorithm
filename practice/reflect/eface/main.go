package main

import "fmt"

func main() {
	var a lift
	a = &worker{name: "aa", age: 17}
	var b goodLift
	//b = &worker{name: "bb", age: 29}
	//b = a
	//b.Eat()
	b = a.(goodLift)
	b.Sleep()

}

type lift interface {
	Eat()
	Work()
	Sleep()
}

type goodLift interface {
	Eat()
	Sleep()
}

type worker struct {
	name string
	age  int
}

func (w *worker) Eat() {
	fmt.Println(w.name + "eat")
}

func (w *worker) Work() {
	fmt.Println(w.name + "work")
}

func (w *worker) Sleep() {
	fmt.Println(w.name + "Sleep")
}
