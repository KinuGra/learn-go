package main

import "fmt"

func main() {
	const (
		a = 2
		b = "hello"
	)
	fmt.Println(a, b)
	fmt.Println(a, "", b)
	fmt.Printf("%v%v\n", a, b)
}
