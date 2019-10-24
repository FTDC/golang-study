package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//  Vertex 只能是结构体的指针类型
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.Y + v.Y*v.X)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

}

type MyFloat float64

//  http://go-tour-zh.appspot.com/methods/2
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
