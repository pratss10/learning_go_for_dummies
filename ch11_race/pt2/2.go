package main

//trying the wait group function

import (
	"fmt"
	// "time"
	"sync"
)

var balance int

// var mutex = &sync.Mutex{}

func credit(wg *sync.WaitGroup, amount int) {
	// mutex.Lock()
	balance += amount
	// time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Println("Credit:", balance)
	// mutex.Unlock()
	wg.Done()
}

func debit(wg *sync.WaitGroup, amount int) {
	// mutex.Lock()
	balance -= amount
	// time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Println("Debit:", balance)
	// mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	balance = 200

	wg.Add(1)
	go credit(&wg, 100)

	wg.Add(1)
	go debit(&wg, 100)

	wg.Wait()
	fmt.Println("Final balance:", balance)
}
