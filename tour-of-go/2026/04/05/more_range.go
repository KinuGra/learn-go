package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	hoge := []int{33, 22, 11, 55, 88}
	for i, v := range hoge {
		fmt.Println(i, v)
	}
}
