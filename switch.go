package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend:", time.Now().Weekday())
	default:
		fmt.Println("It is a weekday", time.Now().Weekday())
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon:", t)
	default:
		fmt.Println("It's after noon:", t)
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("Unknown type")
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("string")
}
