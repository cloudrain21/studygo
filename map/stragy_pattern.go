// simple strategy pattern
package main

import "fmt"

func main() {
	test()
}

var f map[string]func(int, int) int

func initMap() {
	f = make(map[string]func(int, int) int)
	f["+"] = add
	f["-"] = sub
}

func test() {
	if !testCalculate("+", 1, 2, 3) {
		return
	}
	if !testCalculate("+", 3, 2, 5) {
		return
	}
	if !testCalculate("-", 1, 3, -2) {
		return
	}

	fmt.Println("success")
}

func testCalculate(op string, a, b, want int) bool {
	var r int
	if v, ok := f[op]; ok {
		r = v(a, b)
	}

	if r != want {
		fmt.Printf("fail %d != %d\n", r, want)
		return false
	}
	return true
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
