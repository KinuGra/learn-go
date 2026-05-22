package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	for _, w := range strings.Fields(s) {
		counts[w]++
	}
	return counts
}

func main() {
	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
	// → map[I:2 Then:1 a:1 another:1 ate:2 donut.:2]
}
