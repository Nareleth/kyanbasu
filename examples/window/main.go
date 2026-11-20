package main

import (
    "time"
    "github.com/Nareleth/kyanbasu"
)


func main() {
    // Initialize term
    term := kyanbasu.NewTerm()
    defer term.CloseTerm()

    // Initialize screen
    screen := kyanbasu.NewScreen(term)

    // Init windows
    window1 := kyanbasu.NewWindow(screen, 0,  0, 20, 10)
    window2 := kyanbasu.NewWindow(screen, 22, 0, 20, 10)


    // Loop
    for t := 0; t < 5; t++ {

        if t == 2 {
            window2.Hide()
        }

        // Render Borders
        window1.DrawBorder()
        window2.DrawBorder()

        // Render text
        window1.WriteString(0, 0, "Hello")
        window2.WriteString(0, 0, "World!")

        // Flush
        screen.Flush()

        // End frame
        time.Sleep(1 * time.Second)
    }

}
