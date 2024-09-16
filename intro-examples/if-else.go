package main

import "fmt"

func main() {

	// Simple if-else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Simple if
	if 8%4 == 0 {
		fmt.Println("8 divisible by 4")
	}

	// Multi conditional if-else
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("Either 8 or 7 are even")
	}

	// Declared variable in if-else condition
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
