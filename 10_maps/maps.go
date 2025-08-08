package main

import (
	"fmt"
	"maps"
)

func main() {

	// names := make(map[string]string), // inside square bracket [key] value

	// names["foods"] = "Green Apple"
	// names["area"] = "Backend"

	// fmt.Println(names["foods"], names["area"])

	// fmt.Println(names["phones"])


	// m := make(map[string]int)

	// m["age"] = 30
	// m["price"] = 100

	// fmt.Println(m["books"])
	// fmt.Println(len(m))

	// fmt.Println("Before delete: ")
	// fmt.Println(m)

	// delete(m, "price")

	// fmt.Println("After delete: ")
	// fmt.Println(m)


	// fmt.Println("after clearing the map: ")
	// clear(m)
	// fmt.Println(m)



	// creating map without using the mmake func

	// m := map[string]int{"phones": 0171, "price": 340}

	// fmt.Println(m)


	//checking if item is available in the map
	// value, ok := m["price"]
	// fmt.Println(value)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	m := map[string]int{"phones": 0171, "price": 340}
	m2 := map[string]int{"phones": 0171, "price": 340}

	fmt.Println(maps.Equal(m, m2))

}
