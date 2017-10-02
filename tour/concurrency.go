package main

import (
	"fmt"
	"time"
	"sync"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// Only the sender should close a channel, never the receiver. Sending on a 
	// closed channel will cause a panic.
	// Channels aren't like files; you don't usually need to close them. Closing
	// is only necessary when the receiver must be told there are no more values coming, 
	// such as to terminate a range loop.
	close(c)
}

func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// SafeCounter is safe to use concurrently
type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}
// Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v
	// make sure c.mux will unlock
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	go say("world")
	say ("hello")

	// Channels are a typed conduit through which you
	// can send and receive values with the channel operator, <-.
	// like map and slice, channel must be created before use

	// By default, sends and receives block until the other side is ready. 
	// This allows goroutines to synchronize without explicit locks or condition variables.
	var a = []int{1, 2, 3, 4, 5}

	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x + y)

	// buffered channel: Sends to a buffered channel block only when the buffer is full.
	// Receives block when the buffer is empty.
	bc := make(chan int, 2)
	bc <- 1
	bc <- 2
	fmt.Println(<-bc)
	fmt.Println(<-bc)

	// channel with sender close
	cc := make(chan int, 10)
	go fibonacci(cap(cc), cc)
    //The loop for i := range c receives values from the channel repeatedly until it is closed.
	for i := range cc {
		fmt.Println(i)
	}
	// ok is false if there are no more values to receive and the channel is closed.
	v, ok := <-cc
	fmt.Println(v, ok)

	// select
	sc := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-sc)
		}
		quit <- 0
	}()
	fibonacciSelect(sc, quit)

	// sync.Mutex
	dc := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go dc.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(dc.Value("somekey"))

	// select default
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(1000 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM")
			return
		// Use a default case to try a send or receive without blocking
		default:
			fmt.Println("  ...   ")
			time.Sleep(50 * time.Millisecond)
		}
	}
}