package main

import (
    "fmt"
    "io/ioutil"
    "os/exec"
)

func main() {

    // dateCmd := exec.Command("date")  // for Linux/UNIX
    dateCmd := exec.Command("cmd", "/C", "date", "/T")  // for Windows

    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    // fmt.Println("> date")
    fmt.Println("> date /T")
    fmt.Println(string(dateOut))

    // grepCmd := exec.Command("grep", "hello")  // for Linux/UNIX
    grepCmd := exec.Command("cmd", "/C", "findstr", "hello")  // for Windows

    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()

    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd.Wait()

    // fmt.Println("> grep hello")
    fmt.Println("> findstr hello")
    fmt.Println(string(grepBytes))

    // lsCmd := exec.Command("bash", "-c", "ls -a -l -h")  // for Linux/UNIX with BASH shell
    lsCmd := exec.Command("cmd", "/C", "dir")  // for Windows (should also work in most Linux/UNIX OSes)
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    // fmt.Println("> ls -a -l -h")
    fmt.Println("> dir")
    fmt.Println(string(lsOut))
}