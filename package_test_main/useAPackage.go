package main

import (
	"fmt"
	"github.com/cloudrain21/studygo/package_test"
)

func main() {
	fmt.Println("Using aPackage!")
	package_test.A()
	package_test.B()
	package_test.C()
	package_test.D()
	fmt.Println(package_test.MyConstant)
}
