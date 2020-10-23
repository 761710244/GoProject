package main

import "fmt"

type Action interface {
	run()
	laugh()
}

type People1 struct {
}

type People2 struct {
}

func (Li People1) run() {
	fmt.Println("person1 running")
}

func (Li People1) laugh() {
	fmt.Println("person1 laughing")
}

func (Niu People2) run() {
	fmt.Println("person2 running")
}

func (Niu People2) laugh() {
	fmt.Println("person2 laughing")
}

func main() {
	var action Action
	action = new(People1)
	action.run()
	action.laugh()

	action = new(People2)
	action.laugh()
	action.run()
}
