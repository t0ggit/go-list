package main

import "list/storage/list"

func main() {
	l := list.NewList()
	l.Add(42)
	l.Add(54)
	l.Add(112)
	l.Add(1337)
	l.Print()

	ll := list.NewList()
	ll.Print()
}
