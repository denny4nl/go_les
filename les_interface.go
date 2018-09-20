package main

import (
	"fmt"
)

func PrintAll(val interface{}) {
	fmt.Println(val)
}

func main() {
	names := []string{"stanley", "david", "oscar"}
	for _, n := range names {
		PrintAll(n)
	}
}
