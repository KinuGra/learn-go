package main

import "fmt"

// func fibonacci() func() int {
// 	i := 0
// 	s := []int{0, 1}
// 	return func() int {
// 		if i == 0 {
// 			i++
// 			return 0
// 		} else if i == 1 {
// 			i++
// 			return 1
// 		} else {
// 			var value int = s[len(s)-1] + s[len(s)-2]
// 			s = append(s, value)
// 			return value
// 		}
// 	}
// }

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
