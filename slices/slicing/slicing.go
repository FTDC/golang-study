package main

import "fmt"

func main() {
	p := []int{2, 34, 5, 6, 11, 23}
	fmt.Println("p==", p)
	fmt.Println("p[1:4]", p[1:4])

	fmt.Println("p[:3] ==", p[:3])
	fmt.Println("p[4:] ==", p[4:])
}
