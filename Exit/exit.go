/*
    In Linux/UNIX shell, the command '$?', get exit status or return value of last executed command.
    Read more at:
        https://stackoverflow.com/q/7248031
        https://stackoverflow.com/q/6834487
*/

package main

import (
    "fmt"
    "os"
)

func main() {

    defer fmt.Println("!")  // defers will not run when called os.Exit() function

    os.Exit(3)  // exit with status 3
}