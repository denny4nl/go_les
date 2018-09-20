package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	arr := [...]string{"hello", "world", "denny"}
	fmt.Println("the pointer is", &arr[0])
	var myPointer *string
	fmt.Println("my pointer is", myPointer, &myPointer)
	myPointer = &arr[0]
	fmt.Println("my pointer value is", myPointer)
	mySlice := arr[:]
	fmt.Printf("myslice is %p, %v, %d, %d\n", mySlice, mySlice, len(mySlice), cap(mySlice))
}
