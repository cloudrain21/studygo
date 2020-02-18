package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum float64 = 0.0
	var i int
	var cnt int = 0

	arguments := os.Args
	for i = 1; i < len(arguments); i++ {
		v, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			sum += v
			cnt++
		}
	}

	fmt.Println(sum, i, cnt)
	fmt.Println(sum / float64(cnt))
}
