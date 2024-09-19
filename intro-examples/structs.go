package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 40
	return &p
}

func main() {

	fmt.Println(person{name: "Bob"})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(&person{name: "Anni", age: 25})

	fmt.Println(newPerson("Jon"))

	s := person{name: "Alex", age: 50}
	fmt.Println(s.age)

	sp := &s
	fmt.Println(sp.name)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Sully",
		false,
	}

	fmt.Println(dog)
}
