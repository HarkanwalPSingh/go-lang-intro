package main

import "fmt"

func sum(nums ...int) int {

	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {

	nums := []int{1, 2, 3, 4, 5}

	fmt.Println("Sum of", nums, "is", sum(nums...))
}
