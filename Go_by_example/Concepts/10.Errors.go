package main

import (
	"errors"
	"fmt"
)

// By convention, errors are the last return value and have type error, a built-in interface
func f1(arg int) (int, error) {
	if arg == 42 {
		//errors.New constructs a basic error value with the given error message
		return -1, errors.New("can't work with 42")
	}
	//A nil value in the error position indicates that there was no erro
	return arg + 3, nil
}

// It’s possible to use custom types as errors by implementing the Error() method on them
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

// The two loops below test out each of our error-returning functions
// the use of an inline error check on the if line is a common idiom in Go code
func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
