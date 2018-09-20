package main

import "fmt"
import "time"

func main() {
	messages := make(chan string)
	signals := make(chan bool)
	go func(myMsg chan string) {
		myMsg <- "hello"
	}(messages)

	time.Sleep(3 * time.Second)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	msg := "hi"
	go func(myMsg chan string) {
		<-myMsg
	}(messages)
	time.Sleep(3 * time.Second)
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
