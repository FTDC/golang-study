package main

import "fmt"

func main() {
	var a []int
	printSlice("a", a)

	a = append(a, 2)
	printSlice("a", a)

	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
