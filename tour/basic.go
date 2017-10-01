package main

// import mutiple package
import (
	"fmt"
	"math/rand"
)

// simple function
func add(x int, y int) int {
	return x + y
}

// function return two values
func swap(x, y string) (string, string) {
	return y, x
}

// function with naked return of named values
func split(sum int) (x, y int) {
	// function level variable
	var i = rand.Intn(10);
	x = sum - i
	y = sum / i + 3
	return
}

// package level variale
var (
	// exported
	Name string = "zhangmin"
	// package scope
	z uint64 = 1 << 64 - 1
)

// const variable
const (
	Pi = 3.1415926
	// high pricision const number
	Big = 1 << 100
)

func Basic() {
	fmt.Printf("Type: %T Value: %v\n", Name, Name)

	fmt.Println(add(rand.Intn(10), rand.Intn(10)))

	// variable with initialiers
	var h, w, sum  = "hello", "world", 20
	k := 2
	fmt.Println(k);

	a, b := swap(h, w)
    fmt.Println(a, b)

	fmt.Println(split(sum))
}