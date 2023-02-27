package main

import (
	"fmt"
	"reflect"
)

type lift interface {
	Eat()
	Work()
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

func main() {
	var a int
	typeOfA := reflect.TypeOf(a) //返回的是type的接口对象
	ValueOfA := reflect.ValueOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	fmt.Println(ValueOfA.Type(), ValueOfA.Kind())
}
