// Exercise: Loops and Functions
// As a way to play with functions and loops, implement the square root function using Newton's method.
// 
// Newton's method is to approximate Sqrt(x) by picking a starting point z first, and repeating:
// 
// Hint: Iterate and return the final value of z as the answer:
// 
// z -= (z*z - x) / (2*z)
// To begin with, repeat the calculation 10 times and see how close you get to the answer for various values (1, 2, 3, ...).
// 
// Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small amount). See if that's more or fewer than 10 iterations. How close are you to the math.Sqrt?
// 
// Hint: To declare and initialize a floating point value, give it floating point syntax or use a conversion:

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt nagative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := float64(1)
	for i := 1; i < 10; i++  {
		z -= (z * z - x) / (2 * z)
		fmt.Println("z = ", z)
	}
	return z, nil
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(Sqrt(float64(i)))
		fmt.Println(math.Sqrt(float64(i)))
	}

	v, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(v)
}
