package main

import (
	"fmt"
	"reflect"
)

func make_empty_slice() {
	fmt.Println("make_empty_slice()")
	slice := make([]int, 0)
	slice1 := []int{}

	for index, value := range slice {
		fmt.Println(index, value)
	}
	for index, value := range slice1 {
		fmt.Println(index, value)
	}
}

func chg_slice_index_value() {
	fmt.Println("chg_slice_index_value()")
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice)
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d ", slice[i])
	}
	fmt.Println()

	newSlice := slice[1:3]
	newSlice[1] = 35

	fmt.Println(newSlice)
}

func index_out_of_range() {
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice[3] = 45
}

func append_slice() {
	fmt.Println("append_slice()")
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice = append(newSlice, 60)
	fmt.Println(newSlice)
	fmt.Println(slice)
}

func append_capacity() {
	fmt.Println("append_capacity()")
	slice := []int{10, 20, 30, 40}
	newSlice := append(slice, 50)
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
	fmt.Println(len(newSlice), cap(newSlice))
	fmt.Println(newSlice)
	newSlice[3] = 200
	fmt.Println(slice)
	fmt.Println(newSlice)
}

func append_slice2() {
	fmt.Println("append_slice2()")
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	slice := source[2:3:3]
	slice = append(slice, "Kiwi")
	fmt.Println(source)
	fmt.Println(slice)
}

func append_func() {
	fmt.Println("append_func()")
	s1 := []int{1, 2}
	s2 := []int{3, 4}

	fmt.Printf("%v\n", append(s1, s2...))
	s3 := append(s1, 3, 4, 5, 6, 7)
	fmt.Println(s3)
}

func iterage_slice() {
	fmt.Println("iterate_slice()")
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("index: %d value: %d\n", index, value)
	}

	slice1 := []int{1, 2, 3, 4}
	for _, value := range slice1 {
		fmt.Printf("value : %d\n", value)
	}
}

func multiple_division_slice() {
	fmt.Println("multiple_division_slice()")
	slice := [][]int{{10}, {100, 200}}
	for index, value := range slice {
		fmt.Printf("%d : %s\n", index, reflect.TypeOf(value))
		fmt.Print(" ")
		fmt.Println(value)
	}
	slice[0] = append(slice[0], 20)
	for index, value := range slice {
		fmt.Printf("%d : %s\n", index, reflect.TypeOf(value))
		fmt.Print(" ")
		fmt.Println(value)
	}
}

func foo() []int {
	slice := make([]int, 10)
	slice[2] = 100
	return slice
}

func function_slice() {
	fmt.Println("start function_slice()")
	slice := foo()
	for index, value := range slice {
		fmt.Println(index, value)
	}
	fmt.Println("end function_slice()")
}

func main() {
	make_empty_slice()
	chg_slice_index_value()
	//index_out_of_range()
	append_slice()
	append_capacity()
	append_slice2()
	append_func()
	iterage_slice()
	multiple_division_slice()
	function_slice()
}
