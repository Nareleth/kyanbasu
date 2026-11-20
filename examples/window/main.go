package main

import (
    "time"
    "github.com/Nareleth/kyanbasu"
)


func main() {
    term := kyanbasu.NewTerm()
    defer term.CloseTerm()

    screen := kyanbasu.NewScreen(term)

    window1 := kyanbasu.NewWindow(screen, 0,  0, 20, 10)
    window2 := kyanbasu.NewWindow(screen, 22, 0, 20, 10)


    // Draw in window
    window1.WriteString(0, 0, "Hello")
    window2.WriteString(0, 0, "World!")

    // Loop
    // Temp commenting until i remember to add input ending
    /*
    for {
        screen.Flush()
        time.Sleep(5 * time.Second)
    }
    */

    screen.Flush()
    time.Sleep(5 * time.Second)
}
