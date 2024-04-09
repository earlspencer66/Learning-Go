// used often in idiomatic Go
// for example to return both result and error values from a function
package main

import "fmt"

//The (int, int) in this function signature shows that the function returns 2 ints
func vals() (int, int) {
	return 3, 7
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	//a subset of the returned values, use the blank identifier _
	_, c := vals()
	fmt.Println(c)
}
