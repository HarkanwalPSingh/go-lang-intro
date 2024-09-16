package main

import "fmt"

func multiValue() (int, int) {
	return 8, 9
}

func main() {
	_, res := multiValue()
	fmt.Println("Result is:", res)
}
