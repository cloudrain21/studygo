package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer

	b.Write([]byte("Hello "))

	fmt.Fprintf(&b, "Go in Action")

	io.Copy(os.Stdout, &b)
}
