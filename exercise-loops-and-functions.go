package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	z, zold, diff := 1.0, 1.0, 1.0
	steps := 1
	for diff > 0.01 {
		zold = z
		z -= (z*z - x) / 2*z
		if z > zold {
			diff = z - zold
		} else {
			diff = zold - z
		}
		steps += 1
		// fmt.Println("z=", z, "zold=", zold, "diff=", diff)
	}
	fmt.Println("steps=", steps, "diff=", diff)
	return z;
}

func diff(x, y float64) float64 {
	return x - y
}

func try(x float64) {
	s1 := sqrt(x)
	s2 := math.Sqrt(x)
	fmt.Println("sqrt(x)=", s1)
	fmt.Println("Sqrt(x)=", s2, "diff=", diff(s1,s2))
}

func main() {
	try(2)
	try(42)

	var q *int
	*q = 42
}
