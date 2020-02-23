package main

import "fmt"

func main() {
	var i interface{} = 1
	s, ok := i.(string)
	fmt.Println(s, ok)

	ss, ok := i.(int)
	fmt.Println(ss, ok)
}
