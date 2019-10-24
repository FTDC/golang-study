package main

import "fmt"

func main() {
	p := []int{2, 3, 4, 5, 66, 7, 8}

	fmt.Println(cap(p))

	for i := 0; i < len(p); i++ {
		fmt.Println(p[i])
	}
}
