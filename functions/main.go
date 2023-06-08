package main

import "fmt"

func printSomethink(arg1, arg2 string) {
	fmt.Printf("%v%v\n", arg1, arg2)
}

func printAll(args ...string) {
	if len(args) != 0 {
		for _, item := range args {
			fmt.Printf("%v\n", item)
		}
	} else {
		fmt.Printf("Args is empty\n")
	}
}

func pseudoSwap(arg1, arg2 int) {
	fmt.Printf("BEFOREINFUNC\narg1 = %v; arg2 = %v\n", arg1, arg2)
	tmp := arg1
	arg1 = arg2
	arg2 = tmp
	fmt.Printf("AFTERINFUNC\narg1 = %v; arg2 = %v\n", arg1, arg2)
}

func rightSwap(arg1, arg2 *int) {
	fmt.Printf("BEFOREINFUNC\narg1 = %v; arg2 = %v\n", *arg1, *arg2)
	tmp := *arg1
	*arg1 = *arg2
	*arg2 = tmp
	fmt.Printf("AFTERINFUNC\narg1 = %v; arg2 = %v\n", *arg1, *arg2)
}

func sum2val(val1, val2 int) int {
	return val1 + val2
}

func retSwap(arg1, arg2 int) (int, int) {
	return arg2, arg1
}

func retTemp(temp int) int {
	if temp > 0 {
		return 36
	}
	return 0
}

func testDefer() {
	fmt.Printf("1\n")
	defer fmt.Printf("3\n")
	fmt.Printf("2\n")
	defer fmt.Printf("4\n")
}

func main() {
	printSomethink("Hello", ", funcs")
	printAll("Hello", ",", "funcs")
	printAll()

	mySayHello := []string{"Hello", ",", "funcs"}
	printAll(mySayHello...)

	arg1, arg2 := 1, 3
	fmt.Printf("BEFORE\narg1 = %v; arg2 = %v\n", arg1, arg2)
	pseudoSwap(arg1, arg2)
	fmt.Printf("AFTER\narg1 = %v; arg2 = %v\n", arg1, arg2)

	fmt.Printf("BEFORE\narg1 = %v; arg2 = %v\n", arg1, arg2)
	rightSwap(&arg1, &arg2)
	fmt.Printf("AFTER\narg1 = %v; arg2 = %v\n", arg1, arg2)

	fmt.Printf("Sum 1 + 3 = %v\n", sum2val(1, 3))

	arg1, arg2 = 1, 3
	arg1, arg2 = retSwap(1, 3)
	fmt.Printf("%v %v\n", arg1, arg2)

	fmt.Printf("Optimal temp for alive peoples is %v\n", retTemp(1))
	if retTemp(1) == 0 {
		fmt.Printf("U'r is dead\n")
	} else {
		fmt.Printf("U'r is alive\n")
	}

	testDefer()
}
