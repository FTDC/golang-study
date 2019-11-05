package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	c <- sum // 总和送入到 C 管道
}

func main() {

	a := []int{2, 3, 4, 5, 6, 7, 8}

	c := make(chan int)

	go sum(a[:len(a)/2], c)
	go sum(a[:len(a)/2], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
