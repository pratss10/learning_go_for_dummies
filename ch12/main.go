package main

import (
	"fmt"
	"time"
)

// ---send data into a channel---
func sendData(ch chan string) {
	fmt.Println("Sending a string into channel...")
	time.Sleep(5 * time.Second)
	ch <- "Hello"
}

// ---getting data from the channel---
func getData(ch chan string) {
	fmt.Println("String retrieved from channel:", <-ch)
}
func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)
	fmt.Scanln()
}
