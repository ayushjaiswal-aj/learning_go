package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	cV1 := 1.5
	cV2 := int(cV1)
	fmt.Println("1.5 converted to int: ", cV2)

	cV3 := "123"
	cV4, err := strconv.Atoi(cV3)

	if err == nil {
		fmt.Println("123 converted to int: ", cV4)
		fmt.Println("Type of cV4: ", reflect.TypeOf(cV4))
	}

	cV5 := 5000
	cV6 := strconv.Itoa(cV5)
	fmt.Println(cV5, "converted to string: ", cV6)
	fmt.Println("Type of cV6: ", reflect.TypeOf(cV6))

	cV7 := "3.14"
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		fmt.Println("3.14 converted to float64: ", cV8)
		fmt.Println("Type of cV8: ", cV8, reflect.TypeOf(cV8))
	}

}
