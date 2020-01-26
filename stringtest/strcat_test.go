package strcat

import (
	"fmt"
	"strings"
)

func Example_strCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
	// Output:
	// abcdef
	// abcdef
}

func Example_variousStrCat() {
	s := "abc"
	s1 := fmt.Sprint(s, "def")
	s2 := fmt.Sprintf("%sdef", s)
	s3 := strings.Join([]string{s, "def"}, "")
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	// Output:
	// abcdef
	// abcdef
	// abcdef
}
