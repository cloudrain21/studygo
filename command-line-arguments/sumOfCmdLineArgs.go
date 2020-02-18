package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var sum int64 = 0
	arguments := os.Args
	for i := 1; i < len(arguments); i++ {
		v, err := strconv.ParseInt(arguments[i], 10, 32)
		if err == nil {
			sum += v
		}
	}

	fmt.Println(sum)
}
