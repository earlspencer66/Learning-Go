// go's associative data type -- also called hashes or dicts in other languages
package main

import (
	"fmt"
)

func main() {
	//To create an empty map, use the builtin make: make(map[key-type]val-type)
	m := make(map[string]int)

	//Set key/value pairs using typical name[key] = val syntax.
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	// Get a value for a key with name[key].
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	//If the key doesn’t exist, the zero value of the value type is returned
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	//The builtin len returns the number of key/value pairs when called on a map
	fmt.Println("len:", len(m))

	//The builtin delete removes key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// To remove all key/value pairs from a map, use the clear builtin.
	clear(m)
	fmt.Println("map:", m)
}
