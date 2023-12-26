package main

import "fmt"

var pl = fmt.Println

func main() {
	slice1 := make([]string, 6)
	slice1[0] = "Society"
	slice1[1] = "Science"
	slice1[2] = "Technology"
	slice1[3] = "Mathematics"
	slice1[4] = "Physics"
	slice1[5] = "Chemistry"

	pl("slice1 size: ", len(slice1))

	for i := 0; i < len(slice1); i++ {
		pl(slice1[i])
	}

	for _, x := range slice1 {
		pl(x)
	}

	slice2 := []int{12, 24, 36, 48, 60, 72, 84, 96, 108}
	pl("slice2 size: ", len(slice2))

	for i := 0; i < len(slice2); i++ {
		pl(slice2[i])
	}

	for _, x := range slice2 {
		pl(x)
	}

	sliceArray := [5]int{1, 2, 3, 4, 5}

	sl3 := sliceArray[0:2]
	pl("sl3 size: ", len(sl3))
	pl("1st 3 in sl3: ", sl3)

}
