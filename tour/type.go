package main

import (
	"fmt"
	"math"
)

// struct: a collection of fields
type Vertex struct {
	X,Y int
	desp string
}

func main() {
	var i int = 2;
	// ponters like c, but no pointer arithmetic
	p := &i
	fmt.Printf("Type %T, value: %v, value: %v\n", p, p, *p)

	// struct
	v := Vertex{X:1}
	v = Vertex{1, 2, "first"}
	v.X = 1e9
	
	pv := &v
	pv.Y = 1e2
	fmt.Println(*pv)

	// arrays, length is fixed
	var ai [10]int
	fmt.Println(ai[i])

	// slice:flexible view of a array, dynamic length, like reference to arrays
	// slice has length(current slice length) and capacity(underlying array length)
	var bs []int
	bs = ai[:5]
	s := bs[1:]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	var ns []int
	if ns == nil {
		fmt.Println("nil slice")
	}

	// make: create slice, dynamically-sized aray
	var ms []int = make([]int, 2, 5)
	printSlice(ms)

	ams := append(ms, 1, 2, 3, 4,5)
	printSlice(ams)

	// range: iterate over map or slice
	var demo = []int{1, 2, 3, 4}
	for index, value := range(demo) {
		fmt.Println(index, value)
	}
	for _, value := range(demo) {
		fmt.Println(value)
	} 
	for index := range(demo) {
		fmt.Println(index)
	} 

	// map: map key to values
	var m map[string]Vertex
	m = make(map[string]Vertex)
	m["first"] = Vertex{1, 2, "first"}
	m["second"] = Vertex{X: 2}
	fmt.Println(m)
	delete(m, "second")
	elem, ok := m["first"]
	fmt.Println(elem, ok)
	fmt.Println(m)

	// function, like normal variables
	hypot := func (x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// closures
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// function may be closures. A closure is a function value that references variables
// from outside its body. The function may access and assign to the referenced variables;
// in this sense the function is "bound" to the variables.

// For example, the adder function returns a closure. Each closure is bound to 
// its own sum variable.
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x;
		return sum
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}