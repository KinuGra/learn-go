package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// スライスは配列への参照のようなもの
	var s []int = primes[1:4]
	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// defaults
	s2 := []int{2, 3, 5, 7, 11, 13}
	s2 = s2[1:4]
	fmt.Println(s2)

	s2 = s2[:2]
	fmt.Println(s2)

	s2 = s2[1:]
	fmt.Println(s2)
}
