package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(s[:])
	s = s[1:4]
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)

	// スライスはlength, capacityの両方を持つ
	s2 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s2)
	s2 = s2[:0]
	printSlice(s2)
	s2 = s2[:4]
	printSlice(s2)
	s2 = s2[2:]
	printSlice(s2)
	s2 = []int{1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(s2)
}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d %v\n", len(s), cap(s), s)
}
