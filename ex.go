package example

import "fmt"

func Sum(a, b int) int {
	return a + b
}

func Diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func PowerInt(a, b int) int {
	if b == 0 {
		return 1
	}
	ret := a
	for i := 1; i < b; i++ {
		ret *= a
	}
	return ret
}

func Hello() {
	fmt.Println("hello world.")
}
