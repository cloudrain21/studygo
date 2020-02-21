package main

import (
	"bytes"
	"strings"

	//"bytes"
	"fmt"
	//"strings"
	"time"
)

func main() {

	str := ""
	str1 := "AAAAAAAAAA"

	start := time.Now()
	for i := 0; i < 100000; i++ {
		str += str1
	}
	elapsed := time.Since(start)
	fmt.Printf("strlen(%d) : %v\n", len(str), elapsed)

	var b bytes.Buffer
	str = ""
	start = time.Now()
	for i := 0; i < 100000; i++ {
		b.WriteString(str1)
	}
	str = b.String()
	elapsed = time.Since(start)
	fmt.Printf("strlen(%d) : %v\n", len(str), elapsed)

	str = ""
	start = time.Now()
	for i := 0; i < 100000; i++ {
		str = fmt.Sprintf("%s%s", str, str1)
	}
	elapsed = time.Since(start)
	fmt.Printf("strlen(%d) : %v\n", len(str), elapsed)

	str = ""
	mySlice := []string{}
	for i := 0; i < 100000; i++ {
		mySlice = append(mySlice, str1)
	}
	start = time.Now()
	str = strings.Join(mySlice, "")
	elapsed = time.Since(start)
	fmt.Printf("strlen(%d) : %v\n", len(str), elapsed)

	str = ""
	start = time.Now()
	str = strings.Repeat(str1, 100000)
	elapsed = time.Since(start)
	fmt.Printf("strlen(%d) : %v\n", len(str), elapsed)

}
