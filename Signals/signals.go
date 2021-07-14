package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {

    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    // registers the given channel to receive notifications of the specified signals
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()  // goroutine executes a blocking receive for signals

    fmt.Println("awating signal")
    <-done
    fmt.Println("exiting")
}