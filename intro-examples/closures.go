package main

import "fmt"

func intIter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	intSeq := intIter()

	fmt.Println(intSeq())
	fmt.Println(intSeq())
	fmt.Println(intSeq())

	newIntSeq := intIter()
	fmt.Println(newIntSeq())
}
