package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func quotient (a float64, b float64) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return int(a / b), nil
}

func variadicAdd(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	addition := add(1, 2)
	fmt.Printf("addition of 1 and 2 = %d", addition)

	quotient, err := quotient(5, 2)
	if(err != nil) {
		fmt.Println(err)
	} else {
		fmt.Printf("quotient of 5 and 2 = %d", quotient)
	}

	variadicAddition := variadicAdd(1, 2, 3, 4, 5)
	fmt.Printf("addition of 1, 2, 3, 4, 5 = %d", variadicAddition)

}
