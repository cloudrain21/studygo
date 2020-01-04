package strconv

import (
	"fmt"
	"strconv"
)

func Example_strconv() {

	var i int
	var k int64
	var f float64
	var s string
	var err error

	i, err = strconv.Atoi("350")
	fmt.Println(i, err)

	k, err = strconv.ParseInt("cc7fdd", 16, 32)
	fmt.Println(k, err)

	k, err = strconv.ParseInt("0xcc7fdd", 0, 32)
	fmt.Println(k, err)

	f, err = strconv.ParseFloat("3.14", 64)
	fmt.Println(f, err)

	s = strconv.Itoa(340)
	fmt.Println(s, err)

	s = strconv.FormatInt(13402077, 16)
	fmt.Println(s, err)

	// Output:
	// 350 <nil>
	// 13402077 <nil>
	// 13402077 <nil>
	// 3.14 <nil>
	// 340 <nil>
	// cc7fdd <nil>
}
