package example

import "math"

func Sum(a, b int) int {
	return a + b
}

func Diff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
