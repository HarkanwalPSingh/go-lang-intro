package main

import (
	"fmt"
	"slices"
)

func main() {

	var slice1 []string
	fmt.Println("uninit:", slice1, slice1 == nil, len(slice1) == 0)

	slice1 = make([]string, 3)
	fmt.Println("emp:", slice1, "len:", len(slice1), "cap:", cap(slice1))

	slice1[0] = "a"
	slice1[1] = "b"
	slice1[2] = "c"
	fmt.Println("set:", slice1)
	fmt.Println("get:", slice1[2])

	fmt.Println("len:", len(slice1))

	slice1 = append(slice1, "d", "e", "f")
	fmt.Println("apd:", slice1)

	slice2 := make([]string, len(slice1))
	copy(slice2, slice1)
	fmt.Println("cpy:", slice2)

	//Slicing
	slice3 := slice2[2:]
	fmt.Println("slc1:", slice3)

	slice3 = slice2[:3]
	fmt.Println("slc2:", slice3)

	slice3 = slice2[:]
	fmt.Println("slc3:", slice3)

	slice4 := []string{"a", "b", "c"}
	fmt.Println("dcl:", slice4)

	// Utils
	slice5 := []string{"g", "h", "i"}
	if slices.Equal(slice4, slice5) {
		fmt.Println("slice4 == slice5")
	} else {
		fmt.Println("slice4 != slice5")
	}

	slices.Sort(slice4)

	if _, found := slices.BinarySearch(slice4, "c"); found {
		fmt.Println("'c' found in slice4")
	} else {
		fmt.Println("'c' not found in slice4")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
