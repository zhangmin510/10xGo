package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// TODO optimze iteration times
func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 1; i < 10; i++  {
		z -= (z * z - x) / (2 * z)
		// fmt.Println("z = ", z)
	}
	return z
}

func Flow(){
	// for loop
	sum := 0
	for i := 1; i < 10; i++ {
		sum += i;
	}
	// for = while
	for sum < 100 {
		sum += sum
	}

	// forever
	for {
		if sum > 200 {
			break
		}
		sum += 1
	}
	fmt.Println("sum = ", sum);

	// if with statement, k's scope is only in if/else block
	if k := 2; k < 3 {
		fmt.Println(k)
	} else {
		fmt.Println("worng")
	}
	
	for i := 1; i < 10; i++ {
		t := float64(i)
		fmt.Println(Sqrt(t))
		fmt.Println(math.Sqrt(t))
	}

	// switch
	switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X")
		case "linux":
			fmt.Println("Linux")
		case "windows":
			fallthrough
		default:
			fmt.Printf("%s.", os)
	}

	fmt.Printf("\n")

	// deferred call's arguments are evaluated immediately, 
	// but the function call is not executed until the surrounding function returns
	for i := 0; i < 3; i++ {
		defer fmt.Println("defer call:", i);
	}

	// switch true, replacement for long if-then-else statements
	t := time.Now()
	switch {
		case t.Hour() < 12:
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon.")
		default:
			fmt.Println("Good evening.")
	}

}