package main

import (
    "bufio"
    "fmt"
    "os"
    "github.com/Nareleth/kyanbasu"
)


func main() {
    t := kyanbasu.NewTerm()
    defer t.CloseTerm()

    fmt.Println("Press Enter to quit")
    bufio.NewReader(os.Stdin).ReadString('\n')
}
