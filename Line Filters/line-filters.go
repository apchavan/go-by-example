/*
First manually create a file called 'lines' at temporary location with follwing data (no spaces):
    hello
    filter
then run this program.

To run, use "type" command instead of "cat" on Windows (https://stackoverflow.com/a/60254):
Linux/UNIX Terminal -> cat /tmp/lines | go run line-filters.go
Windows CMD -> PowerShell -Command "type %TEMP%/lines | go run line-filters.go"
Windows PowerShell -> type %TEMP%/lines | go run line-filters.go
*/
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        ucl := strings.ToUpper(scanner.Text())
        fmt.Println(ucl)
    }

    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}