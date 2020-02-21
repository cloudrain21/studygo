package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	s := []byte("Data to write\n")

	f1, err := os.Create("/tmp/f1.txt")
	checkError(err)
	defer f1.Close()
	fmt.Fprintf(f1, string(s))

	f2, err := os.Create("/tmp/f2.txt")
	checkError(err)
	defer f2.Close()
	n, err := f2.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)

	f3, err := os.Create("/tmp/f3.txt")
	checkError(err)
	w := bufio.NewWriter(f3)
	n, err = w.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()

	f4 := "/tmp/f4.txt"
	err = ioutil.WriteFile(f4, s, 0644)
	checkError(err)

	f5, err := os.Create("/tmp/f5.txt")
	checkError(err)
	n, err = io.WriteString(f5, string(s))
	checkError(err)
	fmt.Printf("wrote %d bytes\n", n)
}
