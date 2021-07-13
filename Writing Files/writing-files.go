package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    dat1Path := os.TempDir() + "/dat1"
    dat2Path := os.TempDir() + "/dat2"

    d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile(dat1Path, d1, 0644)
    check(err)

    f, err := os.Create(dat2Path)
    check(err)

    defer f.Close()

    d2 := []byte{115, 111, 109, 101, 10}  // Means -> "some\n"
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    n3, err := f.WriteString("writes\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n3)

    f.Sync()

    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    check(err)
    fmt.Printf("wrote %d bytes\n", n4)

    w.Flush()
}