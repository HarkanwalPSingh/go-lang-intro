package main

import (
	"fmt"
	"time"
)

func main() {

	// Basic switch
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Multi expressions switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend baby")
	default:
		fmt.Println("Shit it's weekday :(")
	}

	// if-else switch way
	currentTime := time.Now()
	switch {
	case currentTime.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// Type Switch
	TypeChecker := func(input interface{}) {
		switch t := input.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know the type %T\n", t)
		}
	}
	TypeChecker(true)
	TypeChecker(1)
	TypeChecker(3e20)
	TypeChecker("Hey")
}
