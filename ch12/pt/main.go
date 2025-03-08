package main

import (
	"fmt"
	"time"
)

func fib(n int, c chan int) {
	a, b := 1, 1
	for i := 0; i < n; i++ {
		c <- a
		a, b = b, a+b
		time.Sleep(2 * time.Second)
	}
	close(c)
}
func counter(n int, c chan int) {
	for i := 0; i < n; i++ {
		c <- i
		time.Sleep(1 * time.Second)
	}
	close(c)
}

//sequentially happening

// func main() {
// 	c1 := make(chan int)
// 	c2 := make(chan int)
// 	go fib(10, c1)     // generate 10 fibo nums
// 	go counter(10, c2) // generate 10 numbers
// 	for i := range c1 {
// 		fmt.Println("fib()", i)
// 	}
// 	for i := range c2 {
// 		fmt.Println("counter()", i)
// 	}
// }

//trying to use select statement to print the value as soon as they are available

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go fib(10, c1)     // generate 10 Fibonacci numbers
	go counter(10, c2) // generate 10 numbers
	c1Closed := false
	c2Closed := false
	for {
		select {
		case n, ok := <-c1:
			if !ok {
				// channel closed and drained
				c1Closed = true
				if c1Closed && c2Closed {
					return
				}
			} else {
				fmt.Println("fib()", n)
			}
		case m, ok := <-c2:
			if !ok {
				// channel closed and drained
				c2Closed = true
				if c1Closed && c2Closed {
					return
				}
			} else {
				fmt.Println("counter()", m)
			}
		}
	}
}
