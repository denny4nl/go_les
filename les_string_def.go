package main

import "fmt"

type People struct {
	Name string
}

func (p *People) String() {
	fmt.Println("hello")
}

func main() {
	p := &People{}
	p.String()
}
