package main

import "time"
import "fmt"

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("c1 func")
		c1 <- "result 1"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("c2 func")
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
