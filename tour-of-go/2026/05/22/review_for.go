package main

import "fmt"

func main() {
	// Goにはwhile, do whileはない
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", i)
	}
	fmt.Println()

	x := 1
	for x < 100 {
		x *= 2
		fmt.Printf("%d, ", x)
	}
	fmt.Println()

	var i int = 0
	for {
		i++
		fmt.Printf("%d, ", i)
		if i == 20 {
			break
		}
	}
	fmt.Println()

	s1 := []int{100, 200, 300, 500}
	for i, v := range s1 {
		fmt.Println(i, v)
	}

	for _, v := range make([]int, 5, 5) {
		fmt.Print(v)
	}
	fmt.Println()

	for i := range [5]string{"hello", "world"} {
		fmt.Printf("%d, ", i)
	}
}
