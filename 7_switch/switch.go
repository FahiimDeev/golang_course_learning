package main

import (
	"fmt"
)

func main() {
	// i := 5

	// switch i{
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// case 4:
	// 	fmt.Println("four")
	// default:
	// 	fmt.Println("other")
	// }

	// multiple conditional switch

	// switch time.Now().Weekday(){
	// case time.Saturday, time.Monday:
	// 	fmt.Println("Today is Saturday or Monday")
	// default:
	// 	fmt.Println("Today is not Saturday or Monday")
	// }

	// type switch

	funcTest := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("it's an integer")
		case float32, float64:
			fmt.Println("it's an float")
		case bool:
			fmt.Println("it's a boolean")
		default:
			fmt.Println("this might be a string or other datatypes")
		}
	}

	funcTest(5.600)

}
