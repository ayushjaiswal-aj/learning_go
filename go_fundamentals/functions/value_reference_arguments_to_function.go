package main

import "fmt"

func pass_by_value(x int) int {
	x += 1
	return x
}

func pass_by_reference(x *int) int {
	*x += 1
	return *x
}

func main() {
	x := 1
	fmt.Println(pass_by_value(x))
	fmt.Println(x)
	fmt.Println(pass_by_reference(&x))
	fmt.Println(x)
}
