package main

import (
	"errors"
	"fmt"
)

const (
	maxStackSize = 4096
)

type Stack struct {
	data []interface{}
	len  int
	cap  int
}

var (
	ErrStackFull = errors.New("stack: stack is full")
	ErrNoData    = errors.New("stack: stack has no data")
)

func (s *Stack) Push(v interface{}) error {
	if s.len == s.cap {
		return ErrStackFull
	}

	s.data = append(s.data, v)
	s.len++
	return nil
}

func (s *Stack) Pop() (interface{}, error) {
	if s.len == 0 {
		return nil, ErrNoData
	}

	var v interface{}
	v = s.data[s.len-1]
	s.len--
	return v, nil
}

func (s *Stack) PrintAll() {
	for _, v := range s.data {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func (s *Stack) Size() int {
	return s.len
}

func (s *Stack) Cap() int {
	return s.cap
}

func (s *Stack) SetMaxCap(max int) {
	s.cap = max
}

func main() {
	s := "aabbbaccddeeedcc"
	//s := []int{1,2,3,4,5,6,7,8,9}

	var st Stack
	st.SetMaxCap(maxStackSize)
	fmt.Println(st.Size(), st.Cap())

	for _, v := range s {
		st.Push(v)
	}
	fmt.Println(st.Size())
	st.PrintAll()

	len := st.Size()
	for i := 0; i <= len; i++ {
		v, err := st.Pop()
		if err != nil {
			fmt.Printf("\n%s", err)
			break
		}
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println(st.Size())
}
