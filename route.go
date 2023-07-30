package main

import (
    "fmt"
    "time"
)

func printMessage(message string) {
    for i := 0; i < 5; i++ {
        fmt.Println(message)
        time.Sleep(time.Second) // Introducing a small delay to demonstrate concurrency
    }
}

func main() {
    // Creating a goroutine to execute printMessage concurrently
    go printMessage("Hello from Goroutine!")

    // Continue with the main execution in the current goroutine
    for i := 0; i < 5; i++ {
        fmt.Println("Hello from Main Routine!")
        time.Sleep(time.Second)
    }
}