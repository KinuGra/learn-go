package main

import "fmt"

func add(x, y float32) float32 {
	return x + y
}

func main() {
	fmt.Println("abc")
	fmt.Printf("%f\n", add(100, 50))
}
