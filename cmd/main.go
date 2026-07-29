package main

import (
	"fmt"
	arrays "test/tutorials/Arrays"
)

func main() {
	history1 := [][]string{{"milk", "bread"}, {"apple", "banana"}}
	history2 := [][]string{{"milk", "bread"}, {"apple", "banana"}}

	result := arrays.AreOrderHistoriesEqual(history1, history2)

	fmt.Println(result)
}
