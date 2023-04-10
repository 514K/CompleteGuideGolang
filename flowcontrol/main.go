package main

import "fmt"

func main() {
target:
	ok := "ok"

	switch {
	case ok == "ok":
		fmt.Printf("ok\n")
	default:
		fmt.Printf("not ok\n")
	}

	switch test := true; test {
	case test:
		fmt.Printf("test is true\n")
		fallthrough
	case !test:
		fmt.Printf("test is false\n")
		// fallthrough
	default:
		fmt.Printf("is default\n")
		goto target
	}
}
