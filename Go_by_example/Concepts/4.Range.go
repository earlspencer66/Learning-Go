// range iterates over elements in a variety of data structures.
package main

import "fmt"

func main() {
	//Here we use range to sum the numbers in a slice.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//range on arrays and slices provides both the index and value for each entry.
	//above we didn’t need the index, so we ignored it with the blank identifier _
	for i, num := range nums {
		if num == 4 {
			fmt.Println("index:", i)
		}
	}
	// range on map iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	//range on strings iterates over Unicode code points.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
