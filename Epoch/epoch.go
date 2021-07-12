package main

import (
    "fmt"
    "time"
)

func main() {

    now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)

    millis := nanos / 1000000  // 1 millisecond = 1,000,000 nanoseconds (or 1e+6 nanoseconds)
    fmt.Println(secs)
    fmt.Println(millis)
    fmt.Println(nanos)

    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
}