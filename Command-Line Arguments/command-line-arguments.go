/*
    To experiment with command-line arguments it’s best to build a binary with go build first.
    Linux/UNIX ->
        -> go build command-line-arguments.go
        -> ./command-line-arguments a b c d
    Windows ->
        -> go build command-line-arguments.go
        -> command-line-arguments.exe a b c d
*/
package main

import (
    "fmt"
    "os"
)

func main() {

    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}