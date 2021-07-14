/*
    This program may not work on Windows OS due to the issue:
        https://github.com/golang/go/issues/30662
*/
package main

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
)

func main() {

    binary, lookErr := exec.LookPath("ls")  // for Linux/UNIX
    if lookErr != nil {
        panic(lookErr)
    }

    fmt.Println(binary)
    args := []string{"ls", "-a", "-l", "-h"}  // for Linux/UNIX

    env := os.Environ()

    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}