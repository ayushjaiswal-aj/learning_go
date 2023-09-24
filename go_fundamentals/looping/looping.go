package main

import "fmt"

var pl = fmt.Println

func main() {
	for x := 1; x < 10; x++ {
		pl("x = ", x)
	}

	for x := 10; x > 0; x-- {
		pl("x = ", x)
	}

	x := 0
	for true {
		x++
		if x == 10 {
			break
		}
		pl("x = ", x)
	}
}
