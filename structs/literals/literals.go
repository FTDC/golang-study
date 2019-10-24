package main

import "fmt"

type Vertex struct {
	x, y int
}

var (
	v1 = Vertex{1, 3}
	v2 = Vertex{x: 1}
	v3 = Vertex{y: 6}
	v4 = Vertex{}
	p  = &Vertex{1, 2}
	p2 = &v1
)

func main() {
	fmt.Println(v1, p, p2, v2, v3, v4)
}
