package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := []int{}
	sliceSize := 10
	for i := 0; i < sliceSize; i++ {
		s = append(s, rand.Intn(100))
	}
	c := make(chan int)
	partSize := 2
	parts := sliceSize / partSize
	i := 0
	for i < parts {
		go sum(s[i*partSize:(i+1)*partSize], c)
		i += 1
	}
	i = 0
	total := 0
	for i < parts {
		partialSum := <-c // read from channel
		fmt.Println("Partial Sum: ", partialSum)
		total += partialSum
		fmt.Println("Total Sum: ", total)
		i += 1
	}
	fmt.Println("Total: ", total)
	fmt.Scanln()
}
