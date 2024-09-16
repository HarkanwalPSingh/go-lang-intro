package main

import "fmt"

func main() {

	// Loop 1
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// Loop 2
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// Loop 3
	for i := range 3 {
		fmt.Println("range", i)
	}

	// Loop 4
	for {
		fmt.Println("infinite breaked loop")
		break
	}

	// Loop 5
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("Odd number:", n)
	}
}
