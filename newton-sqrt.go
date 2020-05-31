package main

import (
	"fmt"
	"math"
)

const Epsilon = 1e-10

func Sqrt(x float64) float64 {
	z := 1.0
	change := 1.0
	for change > Epsilon {
		new_z := z - ((z*z - x) / (2 * x))
		change = math.Abs(new_z - z)
		z = new_z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(25))
}
