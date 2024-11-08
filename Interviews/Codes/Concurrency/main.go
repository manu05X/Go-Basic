package main

import (
	"fmt"
	"time"
)

func printNumber(num int) {
    fmt.Println(num)
}

func main() {
    for i := 1; i <= 5; i++ {
        go printNumber(i)  // Creates a goroutine
    }
    time.Sleep(1 * time.Second) // Allows goroutines to finish execution
}
