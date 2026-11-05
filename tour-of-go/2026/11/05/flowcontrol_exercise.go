package main

import "fmt"

var z float64 = 1

func Sqrt(x float64) float64 {
	z -= (z*z - x) / (2 * z)
	return z
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(Sqrt(2))
	}
}
