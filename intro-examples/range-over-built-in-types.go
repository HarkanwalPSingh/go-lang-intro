package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 2, 3, 4, 5}
	numMap := map[string]int{"a": 1, "b": 2, "c": 3}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	sum = 0
	for i, num := range nums {
		fmt.Println("value at index", i, "is", num)
		sum += num
	}
	fmt.Println(sum)

	for key, value := range numMap {
		fmt.Printf("%s -> %d\n", key, value)
	}

	keys := []string{}

	for key := range numMap {
		keys = append(keys, key)
	}
	fmt.Println("The map contains the following keys", keys)
}
