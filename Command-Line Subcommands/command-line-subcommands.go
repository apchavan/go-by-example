/*
    To experiment with command-line subcommands it’s best to build a binary with go build first.
    Linux/UNIX ->
        -> go command-line-subcommands.go
        -> ./command-line-subcommands foo -enable -name=joe a1 a2
        -> ./command-line-subcommands bar -level 8 a1
        -> ./command-line-subcommands bar -enable a1
    Windows ->
        -> go command-line-subcommands.go
        -> command-line-subcommands.exe foo -enable -name=joe a1 a2
        -> command-line-subcommands.exe bar -level 8 a1
        -> command-line-subcommands.exe bar -enable a1
*/
package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {

    fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
    fooEnable := fooCmd.Bool("enable", false, "enable")
    fooName := fooCmd.String("name", "", "name")

    barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
    barLevel := barCmd.Int("level", 0, "level")

    if len(os.Args) < 2 {
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }

    switch os.Args[1] {

    case "foo":
        fooCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'foo'")
        fmt.Println("  enable:", *fooEnable)
        fmt.Println("  name:", *fooName)
        fmt.Println("  tail:", fooCmd.Args())
    case "bar":
        barCmd.Parse(os.Args[2:])
        fmt.Println("subcommand 'bar'")
        fmt.Println("  level:", *barLevel)
        fmt.Println("  tail:", barCmd.Args())
    default:
        fmt.Println("expected 'foo' or 'bar' subcommands")
        os.Exit(1)
    }
}