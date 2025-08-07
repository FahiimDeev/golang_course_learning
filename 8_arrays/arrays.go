package main

import "fmt"

// array is a length
func main() {
	//zeroed values
	// int -> 0, string -> "", bool -> false
	// var number [4]int
	// var flag [4]bool
	// var names [4]string

	//  fmt.Println(len(number))
	// number[0] = 1
	// number[1] =
	// number[2] =
	// number[3] =
	// fmt.Println(number)

	// **
	// flag [2] = true
	// fmt.Println(flag)

	// **
	// names[0] = "golang"
	// fmt.Println(names)

	// nums := [3]int{1,2,3}

	// fmt.Println(nums)

	// 2d array
	number := [2][2]int{{1,2},{3,4}}

	fmt.Println(number)

	// so basically we use arrays when we know the size of the array from the beginning,
	// otherwise we use slices for better, dynamically memory allocation

}
