package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

var fib func(n int) int

func main() {
	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	num := 10
	fmt.Println("Factorial of", num, "is", fact(num))
	fmt.Println("Fibonacci of", num, "is", fib(num))
}
