/*
    Additional resources:
        https://www.12factor.net/config

    Run Commands:
    Linux/UNIX ->
        -> go run environment-variables.go
        -> BAR=2 go run environment-variables.go
    Windows ->
        -> go run environment-variables.go
        -> SET BAR=2 && go run environment-variables.go
*/
package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {

    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    fmt.Println()
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair[0])
    }
}