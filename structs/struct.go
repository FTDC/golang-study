package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {

	v := Vertex{1, 2}
	fmt.Println(v)

	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v)

	//结构体指针
	p := &v
	p.X = 2e9
	fmt.Println(v)
}
