package main

import "fmt"

func main() {
	// C言語とは異なり、ポインタ演算はない
	var p *int
	i := 42
	p = &i

	fmt.Println(*p)
	*p = 100
	fmt.Println(*p)
}
