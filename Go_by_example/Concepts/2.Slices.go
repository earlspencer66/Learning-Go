package main

import (
	"fmt"
)

// Unlike arrays, slices are typed only by the elements they contain (not the number of elements).
func main() {
	//An uninitialized slice equals to nil and has length 0
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	//To create an empty slice with non-zero length, use the builtin make.
	//Here we make a slice of strings of length 3 (initially zero-valued)
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	//set and get
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	//len returns the length of the slice
	fmt.Println("len:", len(s))

	// other operations with slices
	// append, which returns a slice containing one or more new values
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	//Slices can also be copy’d.
	// Here we create an empty slice c of the same length as s and copy into c from s
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a “slice” operator with the syntax slice[low:high]
	// For example, this gets a slice of the elements s[2], s[3], and s[4]
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) s[5]
	l1 := s[:5]
	fmt.Println("sl2:", l1)

	// this slices up from (and including) s[2].
	l2 := s[2:]
	fmt.Println("sl2:", l2)

	// We can declare and initialize a variable for slice in a single line as well
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
}
