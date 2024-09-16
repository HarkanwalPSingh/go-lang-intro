package main

import (
	"fmt"
	"math"
)

const constantString string = "constant string duh!"

func main() {
	fmt.Println(constantString)

	const number = 500000000

	const doubleNumber = 3e20 / number
	fmt.Println(doubleNumber)

	fmt.Println(int64(doubleNumber))

	fmt.Println(math.Sin(number))
}
