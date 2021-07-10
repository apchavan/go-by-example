package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    // WaitGroup must be passed to functions by pointer

    // on return, notify the WaitGroup that we’re done
    defer wg.Done()

    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {

    // WaitGroup is used to wait for all the goroutines launched here to finish
    var wg sync.WaitGroup

    // launch several goroutines and increment the WaitGroup counter for each
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    // block until the WaitGroup counter goes back to 0;
    // all the workers notified they’re done.
    wg.Wait()
}