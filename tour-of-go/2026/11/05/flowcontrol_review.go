package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func hoge() string {
	fmt.Println("hello, ")
	return "world"
}

func main() {
	defer fmt.Println(hoge())

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	cnt := 0
	for {
		cnt++
		if cnt > 100 {
			fmt.Println(cnt)
			break
		} else if cnt == 50 {
			fmt.Println(cnt)
		}
	}

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(4, 3, 60),
	)

	// switch
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println(("OS X."))
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today, time.Now(), time.Saturday, today+1)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tommorow.")
	case today + 2:
		fmt.Println("In two days")
	default:
		fmt.Println("Too far away.")
	}

	// switch true
	t := time.Now()
	fmt.Printf("%T\n", t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// defer
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

// if
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
