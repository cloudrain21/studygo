package main

import "fmt"

func f(array [4]int, slice []int) {
	fmt.Println("-----")
	array[0] = 11
	array[1] = 22
	slice[0] = 11
	slice[1] = 22
	fmt.Println("-----")
}

func main() {
	array := [4]int{1, 2, 3, 4}
	slice := []int{1, 2, 3, 4}

	f(array, slice)

	fmt.Print(array, " ")
	fmt.Println()
	fmt.Print(slice, " ")
}
