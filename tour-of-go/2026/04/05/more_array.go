package main

import "fmt"

func main() {
	var a [10]int
	i := 0
	for true {
		a[i] = i

		if n := a[i] * 2; n == 4 {
			fmt.Println("n == 4")
		}

		if i == 9 {
			break
		}

		i++
	}

	var b [2]string
	b[0] = "Hello"
	b[1] = "World"
	fmt.Println(b[0], b[1])
	fmt.Println(b)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
