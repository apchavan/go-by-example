/*
    Run Commands:
    Linux/UNIX ->
        Start in background -> go run http-servers.go &
        Send request -> curl localhost:8090/hello
        To kill server process, use 'kill' command.
    Windows ->
        Start in background -> start /B go run http-servers.go
        Send request -> curl localhost:8090/hello
        To kill server process, use 'Task Manager'.
*/
package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8090", nil)
}