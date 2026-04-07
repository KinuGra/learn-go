package main

import "fmt"

func main() {
	// range
	slice := []int{2, 3, 5, 7, 11}
	for i, v := range slice {
		fmt.Println(i, v)
	}
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
