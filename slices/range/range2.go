package main

import "fmt"

func main() {
	pow := make([]int, 10)

	for i := range pow {
		print(uint(i))

		pow[i] = 1 << uint(i)
	}

	for _, value := range pow {
		fmt.Println(value)
	}
}
