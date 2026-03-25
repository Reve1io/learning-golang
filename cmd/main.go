package main

import (
	"fmt"
	_ "net/http/pprof"
	arrays "test/tutorials/Arrays"
)

func main() {
	prices := []int{100, 50}
	result := arrays.AddSiscount(prices, 100)
	fmt.Println(result)
}
