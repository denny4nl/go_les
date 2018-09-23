package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	Param
}

func main() {
	s := new(Show)
	s.Param = make(map[string]interface{})
	s.Param["RMB"] = 1000

	my_test := make(map[string]string)
	my_test["son"] = "Joe"
	fmt.Printf("%v", my_test)

	var my_test2 map[string]string
	my_test2 = make(map[string]string)
	my_test2["wife"] = "vickie"
	fmt.Printf("%v", my_test2)
}
