package main

import "fmt"

func main() {
    var a = "initial"
    fmt.Println("a =", a)

    var b, c int = 1, 2
    fmt.Println("b =", b, ", c =", c)

    var d = true
    fmt.Println("boolean d =", d)

    var e int
    fmt.Println("default value for 'int' =", e)

    f := "apple"
    fmt.Println("f =", f)
}