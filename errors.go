package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	var end float64
	z := x / 2
	for {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("x=%v; z=%v;\n", x, z)
		if end == roundFloat(z, 5) {
			break
		}
		end = roundFloat(z, 5)
	}
	return end, nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
	var b []byte
	copy(b, "AAAAAAA")

}
