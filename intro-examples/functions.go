package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a int, b int, c int) int {
	return a + b + c
}

func main() {

	fmt.Println(plus(10, 20))

	fmt.Println(plusPlus(10, 20, 30))
}
