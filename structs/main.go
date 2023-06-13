package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Car struct {
	name  string
	model string
	year  int
}

func newCar(name, model string, year int) *Car {
	return &Car{name: name, model: model, year: year}
}

type Avtosalon struct {
	name string
	cars []Car
}

func main() {
	fmt.Printf("Hello, structs!\n")
	logan := Car{}
	logan.name = "Renault"
	logan.model = "Logan"
	logan.year = 2010

	grandVitara := Car{}
	grandVitara.name = "Suzuki"
	grandVitara.model = "Grand vitara"
	grandVitara.year = 2011

	korsGroup := Avtosalon{}
	korsGroup.cars = []Car{logan, grandVitara}
	korsGroup.name = "Kors group"

	fmt.Printf("%v\n", korsGroup.cars)

	var builder strings.Builder
	json.NewEncoder(&builder).Encode(struct {
		Var1 string
		Var2 int
	}{
		Var1: logan.name,
		Var2: logan.year,
	})
	fmt.Printf("%v", builder.String())

	logan2 := logan
	logan2.model = "Logan 2"
	logan2.year = 2015
	fmt.Printf("%v\n", logan)
	fmt.Printf("%v\n", logan2)

	loganRest := &logan
	loganRest.model = "Logan Rest"
	fmt.Printf("%v\n", logan)
	fmt.Printf("%v\n", *loganRest)
	fmt.Printf("%v\n", logan2)

	c4 := newCar("Citroen", "C4", 2013)
	fmt.Printf("%v\n", *c4)

	anonStruct := new(struct{ sas int })
	anonStruct.sas = 10
	fmt.Printf("%v\n", *anonStruct)
}
