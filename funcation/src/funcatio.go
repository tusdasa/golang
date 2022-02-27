package main

import (
	"fmt"
)

func foo(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return b
}

func foo1(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return b, -1
}

func foo2(a string, b int) (r1 int, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r1 = b
	r2 = b
	return
}

func foo3(a string, b int) (r1, r2 int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	r1 = b
	r2 = b
	return
}

func main() {
	c := foo("a", 10)
	fmt.Println("c =", c)
	d, f := foo1("a", 5000)
	fmt.Println("d=%t,f=%t", d, f)
	d, f = foo2("a", 299)
	fmt.Println("d，f=", d, f)
}
