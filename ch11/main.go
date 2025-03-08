package main

import (
	"fmt"
	"time"
)

func say(s string, times int) {
	for i := 0; i < times; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("DECIDE", 2)
	go say("DAILY", 2)
	fmt.Scanln()
}
