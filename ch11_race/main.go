package main

import (
	"fmt"
	// "math/rand"
	// "time"
	"sync"
)

var balance int
var mutex = &sync.Mutex{} // LOCKING THE SHARED RESOURCE

func credit(amount int) {
	mutex.Lock()
	balance += amount
	// time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Println("Credit:", balance)
	mutex.Unlock()
}

func debit(amount int) {
	mutex.Lock()
	balance -= amount
	// time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Println("Debit:", balance)
	mutex.Unlock()
}

func main() {
	balance = 200
	fmt.Println("Starting balance:", balance)
	go credit(100)
	go debit(100)
	go credit(100)
	go debit(100)
	go credit(100)
	go debit(100)
	go credit(100)
	go debit(100)
	go credit(100)
	go debit(100)
	go credit(100)
	go debit(100)
	go fmt.Println("Final balance:", balance)
	fmt.Scanln()
}

// another way can be to import sync/atomic and use the addint64 function
