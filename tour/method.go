package main

import (
	"fmt"
	"math"
	"io"
	"strings"
	"image"
)

// interface is a set of method signatures
type Abser interface {
	Abs() float64
}
type Vertex struct {
	X, Y float64
}

// ubiquitous Stringer interface
func (v Vertex) String() string {
	return fmt.Sprintf("<%v, %v>", v.X, v.Y)
}

func (v Vertex) Error() string {
	return fmt.Sprintf("error: <%v, %v>", v.X, v.Y)
}

// method: function have a receiver argument between keyword func and func's name
// Type Vertex implements the Interface Abser without using keyword like implements
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// function version
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

// method with a pointer receiver
func (v *Vertex) PAbs(a float64) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y + a)
}

func main() {
	v := Vertex{1, 2}

	// ubiquitous Stringer interface
	fmt.Println("Vertex string:", v)

	fmt.Println(v.Abs())

	var p *Vertex = &v
	// the method call p.Abs() is interpreted as (*p).Abs().
	fmt.Println(p.Abs())
	fmt.Println((*p).Abs())

	fmt.Println(Abs(v))

	fmt.Println(v.PAbs(1.1))
	// Go interprets the statement v.PAbs(1.1) as (&v).Scale(1.1) since the PAbs 
	// method has a pointer receiver.
	fmt.Println((&v).PAbs(1.1)) 

	var abs Abser
	abs = v
	abs = p
	fmt.Println(abs.Abs())

	// interface type's value is nil, no error in go
	var vv Vertex
	// vv is nil, abs hold concrete type value nil, but abs itself is non-nil
	abs = vv
	abs = &vv
	fmt.Println(abs)
	fmt.Println(abs.Abs())
	describe(abs)

	// var absNil Abser
	// runtime error
	// Calling a method on a nil interface is a run-time error because there 
	// is no type inside the interface tuple to indicate which concrete method to call.
	// fmt.Println(absNil.Abs())

	// empty interface
	// An empty interface may hold values of any type. 
	// (Every type implements at least zero methods.)
	// Empty interfaces are used by code that handles values of unknown type. 
	var i interface{}
	describeEmpty(i)

	i = 42
	describeEmpty(i)

	i = "hello"
	describeEmpty(i)

	// A type assertion provides access to an interface value's underlying concrete value.
	// This statement asserts that the interface value i holds the concrete type string  and 
	// assigns the underlying string value to the variable s.
	s := i.(string)
	fmt.Println(s)

	// panic
	// f := i.(float64)
	// fmt.Println(f)

	// test like map's test
	s, ok := i.(string)
	fmt.Println(s, ok)

	// no panic
	g, ok := i.(float64)
	fmt.Println(g, ok)

	// A type switch is a construct that permits several type assertions in series.
	switch v := i.(type) {
	case string:
		fmt.Println("type string, ", v)
	case float64:
		fmt.Println("type float64, ", v)
	default:
		fmt.Println("type ", v)
	}
	
	// The error type is a built-in interface similar to fmt.Stringer:
	// As with fmt.Stringer, the fmt package looks for the error interface when printing values
	if err := v.vertexError(); err != nil {
		fmt.Println(err)
	}

	// Read interface
	r := strings.NewReader("hello reader")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v, err =%v, b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	// Image interface
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())



}

func (v Vertex) vertexError() error {
	return &Vertex{-1, -1}
}

func describe(i Abser) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

