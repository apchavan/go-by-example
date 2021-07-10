package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {
    startTime := time.Now()

    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // start 3 workers without any job hence blocked
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // send 5 jobs & close channel to indicate that's all work we have
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    // collect results of the work; also ensures that the worker goroutines have finished
    for a := 1; a <= numJobs; a++ {
        <-results
    }

    endTime := time.Now()
    fmt.Println("\nreal\t", endTime.Sub(startTime))
}