package main

import (
	"container/list"
	"fmt"
)

func printList(l *list.List) {
	//for t := l.Back(); t != nil; t = t.Prev() {
	//	fmt.Print(t.Value, " ")
	//}
	//fmt.Println()
	fmt.Print("(")
	for t := l.Front(); t != nil; t = t.Next() {
		fmt.Print(t.Value, " ")
	}
	fmt.Print(")")
	fmt.Println()
}

func main() {
	values := list.New()

	e1 := values.PushBack("One")
	e2 := values.PushBack("Two")
	printList(values)
	values.PushFront("Three")
	printList(values)
	values.InsertBefore("Four", e1)
	printList(values)
	values.InsertAfter("Five", e2)
	printList(values)

	values.Remove(e2)
	values.Remove(e2)
	printList(values)

	values.InsertAfter("FiveFive", e2)
	printList(values)
	values.PushBackList(values)
	printList(values)

	values.Init()
	printList(values)

	values.PushBack(e1.Value)
	for i := 0; i < 20; i++ {
		values.PushFront(i)
	}
	printList(values)
}
