package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
)

func main() {
	fmt.Printf("Hello, collections\n")
	arr := new([3]string)
	arr[0] = "Persona 1"
	arr[1] = "Persona 2"
	arr[2] = "Persona 3"

	// error
	// arr[3] = "Persona 4"
	for _, item := range arr {
		fmt.Printf("%v\n", item)
	}

	arr2 := [3]string{"Test", "Simple", "Constructor"}

	for _, item := range arr2 {
		fmt.Printf("%v\n", item)
	}

	multiarr := [3][3][3]int{}
	for i := 0; i < len(multiarr); i++ {
		for j := 0; j < len(multiarr[i]); j++ {
			for k := 0; k < len(multiarr[i][j]); k++ {
				multiarr[i][j][k] = rand.Int() % 10
			}
		}
	}
	fmt.Printf("%v\n", multiarr)

	myslice := []string{"hello", "my", "friend"}
	sort.Strings(myslice)

	for _, item := range myslice {
		fmt.Printf("%v\n", item)
	}

	otherslice := myslice
	fmt.Printf("%v\n", reflect.DeepEqual(otherslice, myslice)) // true

	otherslice[1] = "sas"
	fmt.Printf("%v\n", reflect.DeepEqual(otherslice, myslice)) // same true

	arrayOS := *(*[3]string)(otherslice)
	fmt.Printf("%v\n", arrayOS)

	myMap := make(map[string]int, 3)
	myMap["Alex"] = 22
	myMap["Mike"] = 20
	myMap["John"] = 25

	fmt.Printf("%v\n", myMap)

	for i, item := range myMap {
		fmt.Printf("%v %v\n", i, item)
	}

	if _, exist := myMap["Alex"]; exist {
		fmt.Printf("Item exist\n")
	}

	myMap["Test"] = 13

	if v, exist := myMap["Test"]; exist {
		fmt.Printf("Item %v exist\n", v)
	} else {
		fmt.Printf("Item not exist\n")
	}

	delete(myMap, "Test")

	if v, exist := myMap["Test"]; exist {
		fmt.Printf("Item %v exist\n", v)
	} else {
		fmt.Printf("Item not exist\n")
	}

	for key, val := range myMap {
		fmt.Printf("myMap[\"%v\"] = %v\n", key, val)
	}

	keys := make([]string, 0, len(myMap))
	for key, _ := range myMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("myMap[\"%v\"] = %v\n", key, myMap[key])
	}

	price := "$43.98"
	val := string(price[0])
	sumstr := price[1:]
	sum, errparse := strconv.ParseFloat(sumstr, 64)
	if errparse == nil {
		fmt.Printf("%v in %v\n", sum, val)
	}

	BrokenPrice := []rune("€43.98")
	val = string(BrokenPrice[0])
	sumstr = string(BrokenPrice[1:])
	sum, errparse = strconv.ParseFloat(sumstr, 64)
	if errparse == nil {
		fmt.Printf("%v in %v\n", sum, val)
	}

	BrokenStringPrice := "€43.98"

	for index, symbol := range BrokenStringPrice {
		fmt.Printf("%v %v %v\n", index, symbol, string(symbol))
	}

}
