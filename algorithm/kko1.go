package main

import (
	"fmt"
	"math/rand"
)

func getCount(arr []int) (count int) {
	count = 0
	sum := 0
	for i, v := range arr {
		sum += v
		arr[i] = sum
	}

	for i := 0; i < len(arr)-1; i++ {
		left := arr[i]
		right := sum - left
		if left > right {
			count++
		}
	}

	return
}

func main() {
	a := []int{10, 4, -8, 7}
	fmt.Println(getCount(a))

	a = []int{10, -5, 6}
	fmt.Println(getCount(a))

	a = []int{-3, -2, 10, 20, -30}
	fmt.Println(getCount(a))

	a = make([]int, 100000)
	for i := 0; i < 100000; i++ {
		a[i] = rand.Intn(10000)
	}
	fmt.Println(getCount(a))
}
