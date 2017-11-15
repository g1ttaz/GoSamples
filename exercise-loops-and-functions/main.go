package main

import (
	"fmt"
	"math"
)

func abs(x float64) float64 {
	if x >= 0 {
		return x;
	}
	return -x;
}

func sqrt(x float64) float64 {
	if (x < 0) {
		return math.NaN()
	}
	
	z, zold := 1.0, x
	for abs(z - zold) >= 1.0e-6 {
		zold, z = z, z - (z*z - x) / (2*z)
	}

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
	try(-2)
	try(42)
}
