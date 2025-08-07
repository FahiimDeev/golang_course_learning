package main

import "fmt"

// for -> which is the only loop in golang (construct)
func main(){
	//while loop
	// i := 0

	// for i <= 3{
	// 	fmt.Println(i)
	// 	i++
	// }

	//infinite loop
	// for {
	// 	fmt.Println(1)
	// }

	// classic for loop
	// for i:=0; i<=3; i++{
	// 	if i == 2 {
	// 		continue // continue skipped the 2 and printed the rest
	// 	}
	// 	fmt.Println(i)
	// }

	// range loop -> this is comparatively new feature in golang

	for i := range 4{
		fmt.Println(i)
	}

}