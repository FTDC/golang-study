package main

import "fmt"

func adder() func(int) int {
	sum := 0

	fmt.Println("wer", sum)
	return func(x int) int {
		sum += x
		return sum
	}
}

//Go 函数可以是闭包的。闭包是一个函数值，它来自函数体的外部的变量引用。 函数可以对这个引用值进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。
//
//例如，函数 adder 返回一个闭包。每个闭包都被绑定到其各自的 sum 变量上。

func main() {

	// 方法已存在内存中， sum 成为内存变量
	pos, neg := adder(), adder()

	for i := 0; i < 20; i++ {
		fmt.Println(pos(i), neg(-1*i))
	}
}
