package main

import (
    "fmt"
    "github.com/Nareleth/kyanbasu"
)


func main() {
    t := kyanbasu.NewTerm()
    defer t.CloseTerm()

    w, h := t.GetSize()

    fmt.Printf("Term size: (%d, %d)\r\n", w, h)
    fmt.Printf("Press q to quit\r\n")


    t.Move(10, 5)
    t.SetColor(string(kyanbasu.Green))
    t.WriteByte(kyanbasu.Key_X)
    t.SetColor(string(kyanbasu.Reset))
    t.Flush()
    for {
        key := t.ReadKey()

        switch key {
            case kyanbasu.Key_q, kyanbasu.Key_Q: 
            t.CloseTerm()
            return
        }
    }
}
