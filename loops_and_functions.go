package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z, prev_z := 1.0, 2.0
	for i := 0; math.Abs(prev_z - z) > 0.000001; i++ {
		prev_z = z
		z = z - (math.Pow(z, 2) - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
