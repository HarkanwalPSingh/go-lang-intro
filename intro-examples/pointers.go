package main

import "fmt"

func zeroByVal(a int) {
	a = 0
}

func zeroByReference(aPointer *int) {
	*aPointer = 0
}

func main() {

	a := 1
	zeroByVal(a)
	fmt.Println(a)

	zeroByReference(&a)
	fmt.Println(a)

	fmt.Println(&a)
}
