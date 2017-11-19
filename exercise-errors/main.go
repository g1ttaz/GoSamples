package main

import (
	"fmt"
	"math"
)

// ErrNegativeSqrt is a number
type ErrNegativeSqrt float64

// Error returns an error description
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt returns the square root of a number
// via newton algorithm
func Sqrt(x float64) (float64, error) {
	if x < 0.0 {
		return 0.0, ErrNegativeSqrt(x)
	}

	lastZ, z := x, 1.0

	for math.Abs(z-lastZ) >= 1.0e-6 {
		lastZ, z = z, z-(z*z-x)/(2*z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
