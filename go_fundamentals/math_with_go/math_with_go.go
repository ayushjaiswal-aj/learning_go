package main

import (
	"fmt"
	"math"
)

var pl = fmt.Println

func main() {
	pl("add :: 5 + 4 = ", 5+4)
	pl("Sub :: 5 - 4 = ", 5-4)
	pl("Mul :: 5 * 4 = ", 5*4)
	pl("Div :: 5 / 4 = ", 5/4)
	pl("Mod :: 5 % 4 = ", 5%4)

	iNum := 1
	iNum++
	pl("iNum++ = ", iNum)
	iNum--
	pl("iNum-- = ", iNum)

	pl("Abs(-10) = ", math.Abs(-10))
	pl("Pow(2,3) = ", math.Pow(2, 3))
	pl("Sqrt(4) = ", math.Sqrt(4))
	pl("Cbrt(8) = ", math.Cbrt(8))
	pl("Ceil(4.4) = ", math.Ceil(4.4))
	pl("Floor(4.4) = ", math.Floor(4.4))
	pl("Round(4.4) = ", math.Round(4.4))
	pl("Round(4.6) = ", math.Round(4.6))
	pl("Log2(8) = ", math.Log2(8))
	pl("Log10(100) = ", math.Log10(100))
	pl("Max(5, 4) = ", math.Max(5, 4))
	pl("Min(5, 4) = ", math.Min(5, 4))
	pl("Pi = ", math.Pi)
}
