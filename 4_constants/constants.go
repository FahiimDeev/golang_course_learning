package main

import "fmt"

// const name = "golang"

// age := 30 // by the short hand syntax, it wont work outside of the main function

func main(){
	// const name = "golang"

	// name = "javasript" // this will cause an error because constants cannot be reassigned

	// const age = 30

	// age = 31 // this will also cause an error because constants cannot be reassigned

	// fmt.Println(name)



	const (
		// port = 5000
		// host = "localhost"

		// port := 5000 // the shorthand syntax can't be used for constants
		// host := "localhost"

		port = 5000
		host = "localhost"
	)

	// port = 8080 // constants cannot be reassigned

	fmt.Println(port, host)





}