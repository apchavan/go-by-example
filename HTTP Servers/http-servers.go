/*
    Run Commands:
    Linux/UNIX ->
        Run server in background -> go run http-servers.go &
        Send request -> curl localhost:8090/hello
        To kill background server process, use 'kill' command.
    Windows ->
        Run server in background (https://superuser.com/a/591084) -> start /B go run http-servers.go
        OR run server by opening new CMD (https://superuser.com/a/341603) -> start "" go run http-servers.go
        Send request -> curl localhost:8090/hello
        To kill background server process, use 'Task Manager'.
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