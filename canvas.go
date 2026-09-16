package kyanbasu


import (
    "bufio"
    "fmt"
    "os"
)


/*
    A Canvas is a custom type that represents the full drawing area of the program.
    The canvas holds absolute terminal positions (x, y), sizing (width, height), a string buffer to flush cell contents,
    and a front and back render buffer for rendering performance.
*/


/* Structs */

// Canvas
type Canvas struct {
    X       int
    Y       int
    Width   int
    Height  int
    Writer  *bufio.Writer
    Back    [][]Cell
    Front   [][]Cell
}


/* Functions */

// Initialize a new Canvas
func NewCanvas(x, y, width, height int) *Canvas {
    // Init 2d slices for the render buffer rows
    back    := make([][]Cell, height)
    front   := make([][]Cell, height)

    // Init buffer column slices
    for i := range front {
        back[i]     = make([]Cell, width)
        front[i]    = make([]Cell, width)
    }

    // Return filled canvas
    return &Canvas {
        X:          x,
        Y:          y,
        Width:      width,
        Height:     height,
        Writer:     bufio.NewWriter(os.Stdout),
        Back:       back,
        Front:      front
    }
}


// Clear the terminal
func (c *Canvas) Clear() {
    fmt.Fprint(c.Writer, "\033[2H\033[H")
}


// Flush all rendering and print to screen
func (c *Canvas) Flush() {
    // Iterate through the canvas grid for any buffer changes. If so, change only the new cells.
    for y:= range c.Front {
        for x := range c.Front[y] {
            if c.Front[y][x] != c.Back[y][x] {
                c.Front[y][x] = c.Back[y][x]
                c.Move(x, y)
                c.WriteRune(c.Back[y][x].Char)
            }
        }
    }

    c.Writer.Flush()
}


// Hide the terminal Cursor
func (c *Canvas) HideCursor() {
    fmt.Fprint(c.Writer, "\033[?25l")
}


// Move the terminal cursor to (x, y) coordinate
func (c *Canvas) Move(x, y int) {
    fmt.Fprintf(c.Writer, "\033[%d;%dH", y+1, x+1)
}


// Reveal the terminal cursor
func (c *Canvas) ShowCursor() {
    fmt.Frpint(c.Writer, "\033[?25h")
}

// Write a rune character directly to the terminal cursor location
func (c *Canvas) WriteRune(char rune) {
    fmt.Fprintf(c.Writer, "%c", char)
}

