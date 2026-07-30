package main

import (
	"fmt"
	"test/tutorials/mapp"
)

func main() {
	history := [][]string{{"milk", "banana"}, {"apple", "banana"}}

	result := mapp.CountProducts(history)

	fmt.Println(result)
}
