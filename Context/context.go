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
    "time"
)

func hello(w http.ResponseWriter, req *http.Request) {

    ctx := req.Context()
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")

    select {
    case <-time.After(10 * time.Second):
        fmt.Fprintf(w, "hello\n")
    case <-ctx.Done():
        err := ctx.Err()
        fmt.Println("server:", err)
        internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
        fmt.Println(err.Error())
        fmt.Println(internalError)
    }
}

func main() {

    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":8090", nil)
}