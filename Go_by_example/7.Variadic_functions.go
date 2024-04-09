//Variadic functions can be called with any number of trailing arguments
//For example, fmt.Println is a common variadic function

package main

import "fmt"

// Here’s a function that will take an arbitrary number of ints as arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
}
