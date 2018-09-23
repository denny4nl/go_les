package main

import "fmt"

type student struct {
	Name string
}

func f(v interface{}) {
	switch msg := v.(type) {
	case student:
		fmt.Println(msg.Name, 1)
	case *student:
		fmt.Println(msg.Name, 2)
	}
}

func main() {
	s := new(student)
	s.Name = "huang"
	f(s)
}
