// A goroutine is a lightweight thread of execution
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")
	//To invoke this function in a goroutine, use go f(s)
	//This new goroutine will execute concurrently with the calling one
	go f("goroutine")

	//You can also start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")

	//When we run this program, we see the output of the blocking call first,
	//then the output of the two goroutines.

	//The goroutines’ output may be interleaved,
	//because goroutines are being run concurrently by the Go runtime
}
