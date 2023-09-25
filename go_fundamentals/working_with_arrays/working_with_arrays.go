package main

import "fmt"

var pl = fmt.Println

func main() {
	var arr1 [5]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	arr1[4] = 5
	pl(arr1)

	arr2 := [5]int{1, 2, 3, 4, 5}
	pl(arr2)
	pl("arr2 length: ", len(arr2))
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("arr2[%d] = %d\n", i, arr2[i])
	}

	for i, v := range arr2 {
		fmt.Printf("arr2[%d] = %d\n", i, v)
	}
}
