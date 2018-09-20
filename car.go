package main

import "fmt"

type vehicle interface {
	run()
	load()
}

type car struct {
	name     string
	capacity int
	people   int
}

type truck struct {
	name         string
	max_capacity int
	min_capacity int
}

func (c car) run() {
	fmt.Println("My name is ", c.name)
}

func (c car) load() {
	fmt.Printf("Car can load %d kg weights\n", c.capacity)
}

func (c car) load_people() {
	fmt.Printf("Car can load %d people\n", c.people)
}

func (t truck) run() {
	fmt.Println("My name is ", t.name)
}

func (t truck) load() {
	fmt.Printf("Truck can load %d kg weights at most, %d kg at least\n",
		t.max_capacity, t.min_capacity)
}

func my_action(v vehicle) {
	v.run()
	v.load()
}

func main() {
	ves := []vehicle{car{name: "car", capacity: 100, people: 5},
		truck{name: "truck", max_capacity: 2000, min_capacity: 500}}
	for _, v := range ves {
		my_action(v)
	}
}
