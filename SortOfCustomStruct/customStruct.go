package main

import (
	"fmt"
	"sort"
)

type myInterface interface {
	// Find number of elements in collection
	Len() int

	// Less method is used for identifying which elements among index i and j are lesser and is used for sorting
	Less(int, int) bool

	// Swap method is used for swapping elements with indexes i and j
	Swap(int, int)
}

type customStruct struct {
	name string
	age  int
}

type customStructSortFactor []customStruct

func (structSortFactor customStructSortFactor) Len() int {
	return len(structSortFactor)
}
func (structSortFactor customStructSortFactor) Less(i, j int) bool {
	return structSortFactor[i].age < structSortFactor[j].age
}
func (structSortFactor customStructSortFactor) Swap(i, j int) {
	structSortFactor[i], structSortFactor[j] = structSortFactor[j], structSortFactor[i]
}

func main() {
	stuctList := customStructSortFactor{{name: "Parsa", age: 30}, {name: "Jon", age: 40}, {name: "lili", age: 20}}
	fmt.Println(stuctList)
	sort.Sort(stuctList)
	fmt.Println(stuctList)
}
