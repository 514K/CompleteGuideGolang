package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Hello\n")

	myVal := "101.2"
	myVal2 := "101"

	intVal, _ := strconv.ParseInt(myVal, 0, 0)
	fmt.Printf("%v\n", intVal)

	floatVal, _ := strconv.ParseFloat(myVal, 64)
	fmt.Printf("%v\n", floatVal)

	intVal2, _ := strconv.ParseInt(myVal2, 0, 0)
	fmt.Printf("%v\n", intVal2)

	intVal3, _ := strconv.Atoi(myVal2)
	fmt.Printf("%v\n", intVal3)

	boolVal, _ := strconv.ParseBool(myVal)
	fmt.Printf("%v\n", boolVal)
}
