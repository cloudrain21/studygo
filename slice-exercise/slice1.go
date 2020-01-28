package main

import "fmt"

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

func main() {
	chg_slice_index_value()
	//index_out_of_range()
	append_slice()
}
