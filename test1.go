package main

import (
	"fmt"
)

type Vehicle interface {
	move()
}

type car struct {
	name string
}

type person struct {
	name string
	age  uint8
	contacts
}

type contacts struct {
	email string
	phone string
}

func main() {
	var intCh chan int = make(chan int, 3)
	go factorial(5, intCh)
	strCh := make(chan string, 3)
	strCh <- "Alex"
	fmt.Println(cap(strCh))
	fmt.Println(len(strCh))
	fmt.Println(<-intCh)
	close(intCh)
	close(strCh)
	for i := 0; i < cap(intCh); i++ {
		if val, opened := <-intCh; opened {
			fmt.Println(val)
		} else {
			fmt.Println("Channel closed!")
			break
		}
	}

	results := make(map[int]int)
	structCh := make(chan struct{})
	go factorial2(5, structCh, results)
	<-structCh
	for i, v := range results {
		fmt.Println(i, " - ", v)
	}

	/*var people person = person{"Alexandr", 20, contacts{"alex@gmail.com", "+380990000000"}}
	people.name = "Andrej"
	//var q *person = &people
	//q.input()
	people.print()
	var car1 Vehicle = car{"BMW"}
	car2 := car{"Volvo"}
	car1.move()
	car2.move()*/
}

func (c car) move() {
	fmt.Println(c.name, " goes")
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
	fmt.Println("Age:", l.age)
	fmt.Println("Email", l.email)
	fmt.Println("Phone:", l.phone)
}

func factorial(n int, ch chan int) {

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Println(n, "-", result)

	ch <- result
}

func factorial2(n int, ch chan struct{}, results map[int]int) {
	defer close(ch)
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		results[i] = result
	}
}

func divide(x, y float64) float64 {
	if y == 0 {
		panic("Division by zero!")
	}
	return x / y
}
