package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "ayush"
	var v1, v2 = 1, 3.4
	var v3 = "hello"
	v4 := "world"

	fmt.Println("variable name: " + name)
	fmt.Println("Type of variable name: ", name, "is", reflect.TypeOf(name))
	fmt.Println("Variables v1, v2:", v1, v2)
	fmt.Println("Type of variable v1, v2:", v1, v2, "are", reflect.TypeOf(v1), reflect.TypeOf(v2))
	fmt.Println("Variable v3: " + v3)
	fmt.Println("Type of variable v3: ", v3, "is", reflect.TypeOf(v3))
	fmt.Print("Variable v4: " + v4)
	fmt.Println("Type of variable v4: ", v4, "is", reflect.TypeOf(v4))
}
