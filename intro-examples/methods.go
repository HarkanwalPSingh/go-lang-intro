package main

import "fmt"

type rectangle struct {
	length int
	width  int
}

func (r *rectangle) area() int {
	return r.length * r.width
}

func (r rectangle) perimeter() int {
	return 2 * (r.length + r.width)
}

func main() {
	r := rectangle{
		length: 10,
		width:  20,
	}

	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter", r.perimeter())

	rp := &r
	fmt.Println("Area:", rp.area())
	fmt.Println("Perimeter", rp.perimeter())
}
