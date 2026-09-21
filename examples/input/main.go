package main

import (
    "fmt"
    "os"
    "github.com/Nareleth/kyanbasu"
)


func main() {
    // Set terminal to raw mode
    raw, err := kyanbasu.RawMode(os.Stdin)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    
    // Return terminal to normal after program exit
    defer raw.RawClean()
    
    // Initialize a 1 byte slice for capturing keys
    key := make([]byte, 1)

    w, h := kyanbasu.GetSize()

    canvas := kyanbasu.NewCanvas(0, 0, w, h)

    panel1 := kyanbasu.NewPanel(0, 0, 20, 10, canvas)
    panel1.Style = kyanbasu.Heavy

    panel2 := kyanbasu.NewPanel(0, 11, 20, 3, canvas)
    panel2.Style = kyanbasu.Heavy


    // Clear screen
    canvas.Clear()

    // Draw borders
    panel1.DrawBorders(canvas)
    panel2.DrawBorders(canvas)

    // Print quit instructions
    canvas.Move(0, 14)
    fmt.Fprint(canvas.Writer, "Press q to quit")

    canvas.Flush()

    // Game loop
    for {
        // Read key input
        os.Stdin.Read(key)

        // Print key press event
        panel2.WriteText(1, 1, "You pressed: %c", key[0])

        // Exit program
        if key[0] == kyanbasu.Key_q {
            raw.RawClean()
            os.Exit(0)
        }


        // Flush buffer content and print it
        canvas.Flush()
    }
}
