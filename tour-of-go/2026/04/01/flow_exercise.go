package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)
	fmt.Println(sqrt(n))
	fmt.Println((sqrt2(2)))
}

func sqrt(x float64) float64 {
	value := 1.0
	for {
		if value*value > x {
			value -= 0.1
			return value
		}
		value += 0.0001
	}
}

func sqrt2(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
