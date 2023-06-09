package main

import (
	"fmt"
	"math"
	"math/rand"
)

func myFunc1() {
	fmt.Printf("1\n")
}

func myFunc2() {
	fmt.Printf("2\n")
}

func myFunc3(mF func()) {
	mF()
}

func myFunc4() func() {
	if math.Mod(float64(rand.Int()), 10) > 5 {
		return myFunc1
	}
	return myFunc2
}

type myFuncs func()

func bettermyFunc3(mF myFuncs) {
	mF()
}

func bettermyFunc4() myFuncs {
	if math.Mod(float64(rand.Int()), 10) > 5 {
		return myFunc1
	}
	return myFunc2
}

func retAnonFunc() myFuncs {
	return func() {
		fmt.Printf("Super anon func)\n")
	}
}

var wtf = 1

func currentFunc(val int) func() {
	return func() {
		fmt.Printf("%v\n", val)
	}
}

func actualFunc(val *int) func() {
	return func() {
		fmt.Printf("%v\n", *val)
	}
}

func main() {
	fmt.Printf("Hello, funcs types\n")
	var funcT func()
	if math.Mod(float64(rand.Int()), 10) > 5 {
		funcT = myFunc1
	} else {
		funcT = myFunc2
	}

	funcT()

	myFunc3(funcT)

	myFunc4()()

	var funcTB myFuncs
	if math.Mod(float64(rand.Int()), 10) > 5 {
		funcTB = myFunc1
	} else {
		funcTB = myFunc2
	}

	bettermyFunc3(funcTB)
	bettermyFunc4()()

	anonFunc := func() {
		fmt.Printf("Anon func ;)\n")
	}
	anonFunc()
	retAnonFunc()()

	wtf = 3
	staticVal1 := currentFunc(int(wtf))
	wtf = 2
	staticVal2 := currentFunc(int(wtf))
	staticVal1()
	staticVal2()

	wtf = 3
	actualVal1 := actualFunc(&wtf)
	wtf = 2
	actualVal2 := actualFunc(&wtf)
	actualVal1()
	actualVal2()
}
