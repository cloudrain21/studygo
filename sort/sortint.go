package main

import (
	"fmt"
	"sort"
)

type intSlice []int

func main() {
	a := []int{2, 3, 2, 4, 5, 2, 6, 19, 192, 1, 0}

	sort.Ints(a)
	fmt.Println(a)

	b := []int{2, 3, 2, 4, 5, 2, 6, 19, 192, 1, 0}
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	fmt.Println(b)

	c := []int{2, 3, 2, 4, 5, 2, 6, 19, 192, 1, 0}
	sort.Sort(intSlice(c))
	fmt.Println(c)

	s := []string{"aa", "bs", "we", "aa", "oe", "xx"}
	sort.Strings(s)
	fmt.Println(s)
}

func (s intSlice) Len() int {
	return len(s)
}

func (s intSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s intSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
