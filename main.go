package main

import (
	"fmt"

	"example.com/inventory/inventory"
)

func main() {
	s := inventory.New()
	s.Add("a", 2)
	fmt.Println(s.Total())
}
