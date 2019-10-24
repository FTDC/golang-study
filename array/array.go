package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	//a[2] = "Worlds"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}
