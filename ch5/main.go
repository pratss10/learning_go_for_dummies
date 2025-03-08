package main

import "fmt"

func add(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {
	fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	gen := fib()
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
	fmt.Println(gen())
}
