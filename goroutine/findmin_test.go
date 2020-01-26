package main

import "fmt"

func Min(a []int) int {
	if len(a) == 0 {
		return 0
	}

	min := a[0]
	for _, e := range a[1:] {
		if min > e {
			min = e
		}
	}
	return min
}

func ExampleMin() {
	fmt.Println(Min([]int{
		83, 46, 49, 23, 92,
		48, 39, 91, 44, 99,
		25, 42, 74, 56, 23,
	}))
	// Output: 23
}
