package document

import (
	"fmt"
)

func ExampleGetStrLen() {
	fmt.Println(GetStrLen("I am Rain.i"))
	fmt.Println(GetStrLen(""))
	// Output:
	// 11
	// 0
}

func ExampleGetDouble() {
	fmt.Println(GetDouble(9))
	fmt.Println(GetDouble(1))
	// Output:
	// 81
	// 1
}
