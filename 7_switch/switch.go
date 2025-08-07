package main

import "fmt"


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

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("it's an integer")
		case string:
			fmt.Println("it's a string")
		case bool:
			fmt.Println("it's a boolean")
		default:
			fmt.Println("Others", t)
		}
	}

	whoAmI(6)

}
