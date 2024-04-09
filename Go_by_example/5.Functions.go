package main

import "fmt"

//Here’s a function that takes two ints and returns their sum as an int.
func plus(a int, b int) int {
	return a + b
}
func pluPlus(a, b, c, d int) int {
	return ((a + b) - c) / (d)
}
func main() {
	res := plus(5, 10)
	fmt.Println("5+10 =", res)

	res1 := pluPlus(5, 6, 1, 8)
	fmt.Println("((5+6)-1)/8 = ", res1)
}
