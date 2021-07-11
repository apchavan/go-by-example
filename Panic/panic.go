package main

import "os"

func main() {

    panic("a problem")

    _, err := os.Create(os.TempDir() + "/file")
    if err != nil {
        panic(err)
    }
}