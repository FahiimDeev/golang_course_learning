package main

import "fmt"

func main() {
	// slice -> dynamically declares the memory sizes

	var number []int // uninitialized slice is nil

	fmt.Println(number)
	fmt.Println(number == nil)
	fmt.Println(len(number))

	var number = make([]int, 2)

	//capacity is basically tells us maximum number of elemets can fit
	fmt.Println(cap(number))

	fmt.Println(number == nil)

	var number = make([]int, 0, 5) //we initialy put there 0, so that the 0 number don't come in the output

	number = append(number, 1,2,3,4,5,6,7,8,9,10,11)

	number := []int{}

	number = append(number, 1,2)

	fmt.Println(number)
	fmt.Println(cap(number))
	fmt.Println(len(number))

	s := make([]int, 3)
	s := make([]int, 3)
	s= append(s, 10,20, 30, 40)

	fmt.Println(s)
	fmt.Println(cap(s))
	fmt.Println(len(s))

	// slice with loop

	numbers := make([]int, 3)

	for i := 1; i <= 5; i++ {
		numbers = append(numbers, i)
		fmt.Println("Slice:", numbers, "| Len:", len(numbers), "| Cap:", cap(numbers))
	}

	// copy function

	num1 := []int{1,2,345,435,645,234,578,345}
	num2 := make([]int, len(num1)) // The copy() function in Go only copies up to the length of the destination slice.
	copy(num2, num1)

	fmt.Println(num2)

	number_1 := make([]int, 3)
	number_2 := make([]int, 2)

	number_1[0] = 45
	number_1[1] = 90
	number_1[2] = 120

	copy(number_2, number_1)

	fmt.Println(number_2)
	fmt.Println(cap(number_1))
	fmt.Println(len(number_1))

	// slice operator

	number := []int {1,2,3}

	fmt.Println(number[0:1])
	fmt.Println(number[:2])
	fmt.Println(number[1:])

	slice

	number1 := []int {1,2}
	number2 := []int {1,3}

	// with the equal function, we can compare wheter it's true or false
	fmt.Println(slices.Equal(number1, number2))

	// 2D array using slice

	number := [][]int{{1,2},{3,4}}
	fmt.Println(number)

}
