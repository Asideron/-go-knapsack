package main

import "fmt"

func main() {
	items := []*Item{
		{"Cup", 6, 5},
		{"Headphones", 3, 1},
		{"Phone", 4, 3},
		{"Notebook", 2, 3},
		{"Book", 5, 6},
	}

	backpack := NewBackpack(15)

	for _, item := range items {
		backpack.Add(*item)
	}

	pack, err := backpack.Pack()
	if err != nil {
		fmt.Printf("Pack error: %s", err.Error())
		return
	}

	fmt.Printf("Result value: %d\n\n", pack.Value)

	fmt.Println("Packed items:")
	for _, item := range *pack.Items {
		item.Print()
	}
}
