package main

import "fmt"

type Person struct {
	age    int
	name   string
	number int
}

func main() {
	var Li Person
	Li.age = 25
	Li.name = "Li"
	Li.number = 100
	printPerson(Li)

	var An Person
	An.name = "zhang"
	An.age = 26
	An.number = 200
	printPerson(An)
}

func printPerson(person Person) {
	fmt.Println(person.name)
	fmt.Println(person.age)
	fmt.Println(person.number)
}
