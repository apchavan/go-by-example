package main

import (
    "bytes"
    "fmt"
    "regexp"
)

var pl = fmt.Println

func main() {

    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    pl(match)

    r, _ := regexp.Compile("p([a-z]+)ch")

    pl(r.MatchString("peach"))

    pl(r.FindString("peach punch"))

    pl(r.FindStringIndex("peach punch"))

    pl(r.FindStringSubmatch("peach punch"))

    pl(r.FindStringSubmatchIndex("peach punch"))

    pl(r.FindAllString("peach punch pinch", -1))

    pl(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

    pl(r.FindAllString("peach punch pinch", 2))

    pl(r.Match([]byte("peach")))

    r = regexp.MustCompile("p([a-z]+)ch")
    pl(r)

    pl(r.ReplaceAllString("a peach", "<fruit>"))

    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    pl(string(out))
}