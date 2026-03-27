package main

import (
	"fmt"
	_ "net/http/pprof"
	maps "test/tutorials/Maps"
)

func main() {
	result := maps.CountWords(",Go Go go!")
	fmt.Println(result)
}
