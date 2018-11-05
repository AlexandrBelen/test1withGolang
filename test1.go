package main

import (
	"fmt"
)

type person struct{
	name string
	age uint8
	contacts
}

type contacts struct {
	email string
	phone string
}

func main() {
	var people person=person{"Alexandr",20, contacts{"alex@gmail.com","+380990000000"}}
	people.name="Andrej"
	var q *person = &people
	q.input()
	people.print()

}

func (l *person) input() {
	fmt.Print("Input name: ")
	fmt.Scanf("%s", &l.name)
	fmt.Print("Input age: ")
	fmt.Scanf("%d", &l.age)
	fmt.Print("Input email: ")
	fmt.Scanf("%s", &l.email)
	fmt.Print("Input phone: ")
	fmt.Scanf("%s", &l.phone)
}

func (l person) print() {
	fmt.Println("Name:", l.name)
	fmt.Println("Age:",l.age)
	fmt.Println("Email",l.email)
	fmt.Println("Phone:",l.phone)
}

func factorial(n uint) uint {
	if n >= 0 {
		if n == 0 {
			return 1
		}
		return n * factorial(n-1)
	} else {
		fmt.Println("wrong number")
		return 0
	}

}

func divide(x, y float64) float64 {
	if y == 0 {
		panic("Division by zero!")
	}
	return x / y
}

func changeValue(x *int) {
	*x = (*x) * (*x)
}
