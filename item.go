package main

import "fmt"

type Item struct {
	Name   string
	Weight int
	Value  int
}

func (i *Item) Print() {
	fmt.Printf("i:%s w:%d v:%d\n", i.Name, i.Weight, i.Value)
}
