package main

import "fmt"

func print1() {
	fmt.Println("The first print")
}

func print2() {
	fmt.Println("The second print")
}

func main() {
	fmt.Println("Print in the main")
	defer print1()
	defer print2()
	panic("here is a panic")
}
