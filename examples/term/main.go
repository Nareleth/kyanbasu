// Raw Terminal Demo
// Test the capabilities of everything within the term layer.
// The demo reads screen size, prints instructions, and allows you to
// control a movable cursor.

package main

import (
    "os"
    "time"
    "github.com/Nareleth/kyanbasu"
)


// Config
var FPS = 30
var PColor = kyanbasu.Color_Green


// Type PCursor
// A player controlled movable cursor
type PCursor struct {
    Char    byte
    FG      kyanbasu.Color
    X, Y    int
}


// Initializes a new PCursor
func NewPCursor(char byte, fg kyanbasu.Color, x, y int) *PCursor {
    return &PCursor{
        Char:   char,
        FG:     fg,
        X:      x,
        Y:      y,
    }
}


// Draw the PCursor
func (c *PCursor) Draw(t *kyanbasu.Term) {
    t.Move(c.X, c.Y)
    t.SetColor(c.FG)
    t.WriteByte(c.Char)
    t.SetColor(kyanbasu.Color_Reset)
}

// Move the PCursor
func (c *PCursor) Move(dx, dy int) {
    c.X = c.X + dx
    c.Y = c.Y + dy
}

// Read Key Input
func ReadInput(t *kyanbasu.Term, cursor *PCursor) {
    for {
        // Register a key input
        key := t.ReadKey()

        // Parse key input
        switch key {
            // Close application
            case kyanbasu.Key_q, kyanbasu.Key_Q: 
                t.CloseTerm()
                os.Exit(0)
            // Movement
            case kyanbasu.Key_w:
                cursor.Move(0, -1)
            case kyanbasu.Key_a:
                cursor.Move(-1, 0)
            case kyanbasu.Key_s:
                cursor.Move(0, 1)
            case kyanbasu.Key_d:
                cursor.Move(1, 0)
        }
    }

}

func main() {
    // Initializes a new Term object to control rendering methods.
    // We defer the exit behavior now so it can be captured at all stages.
    t := kyanbasu.NewTerm()
    defer t.CloseTerm()

    // Get the width and height of the terminal.
    w, h := t.GetSize()

    // Create a new Player controllable cursor
    cursor := NewPCursor(kyanbasu.Key_X, PColor, 10, 5)

    // Run goroutine for key input
    go ReadInput(t, cursor)

    // Animation loop
    for {
        // Clear screen every frame (very inefficient but lazy)
        t.Clear()

        // Print user instructions.
        t.Move(0, 0)
        t.WriteString("Term size: (%d, %d)", w, h)
        t.Move(0, 1)
        t.WriteString("Press Q to quit")
        t.Move(0, 2)
        t.WriteString("%-7s %-7s %7s %7s", "W: Up", "A: Left", "S: Down", "D: Right")

        // Draw the PCursor on the screen
        cursor.Draw(t)

        // Flush all terminal drawing instructions.
        t.Flush()
        
        // Wait to run again for fps limit
        time.Sleep(time.Second / time.Duration(FPS))
    }
}
