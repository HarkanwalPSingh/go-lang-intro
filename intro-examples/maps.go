package main

import (
	"fmt"
	"maps"
)

func main() {

	firstMap := make(map[string]int)

	firstMap["k1"] = 7
	firstMap["k2"] = 10

	fmt.Println("map: ", firstMap)

	for _, key := range []string{"k1", "k2", "k3"} {
		fmt.Println(key, "value:", firstMap[key])
	}

	delete(firstMap, "k2")
	fmt.Println("map:", firstMap)

	clear(firstMap)
	fmt.Println("map:", firstMap)

	_, prs := firstMap["k2"]
	fmt.Println("prs:", prs)

	newMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", newMap)

	newMap2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(newMap, newMap2) {
		fmt.Println("newMaps are equal")
	}
}
