/*
    To experiment with command-line arguments it’s best to build a binary with go build first.
    Linux/UNIX ->
        -> go build command-line-arguments.go
        -> ./command-line-flags -word=opt -numb=7 -fork -svar=flag
        -> ./command-line-flags -word=opt
        -> ./command-line-flags -word=opt a1 a2 a3
        -> ./command-line-flags -word=opt a1 a2 a3 -numb=7
        -> ./command-line-flags -h
        -> ./command-line-flags -wat
    Windows ->
        -> go build command-line-arguments.go
        -> command-line-flags.exe -word=opt -numb=7 -fork -svar=flag
        -> command-line-flags.exe -word=opt
        -> command-line-flags.exe -word=opt a1 a2 a3
        -> command-line-flags.exe -word=opt a1 a2 a3 -numb=7
        -> command-line-flags.exe -h
        -> command-line-flags.exe -wat
*/
package main

import (
    "flag"
    "fmt"
)

func main() {

    wordPtr := flag.String("word", "foo", "a string")

    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    flag.Parse()

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}