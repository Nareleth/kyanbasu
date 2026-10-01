package kyanbasu


// The internal terminal driver that uses ANSI escape codes to manipulate the terminal.

import (
    "bufio"
    "fmt"
    "os"
)


type Term struct {
    fd      int
    in      *os.File
    //state
    out     *bufio.Writer
}


// Initialize a new virtual terminal
func NewTerm() *Term {
    fd := int(os.Stdin.Fd())
    
    t := &Term{fd: fd, in: os.Stdin, out: bufio.NewWriter(os.Stdout),}

    t.EnableAlternate()
    // Hide Cursor()
    t.Clear()
    t.Flush()    

    return t
}


// ANSI Clear screen
func (t *Term) Clear() {
    fmt.Fprint(t.out, "\033[2J\033[H")
}


// Close virtual terminal and resets to old state
func (t *Term) CloseTerm() {
    t.DisableAlternate()
    t.Flush()
}


// CSI disable alternate screen buffer
func (t *Term) DisableAlternate() {
    fmt.Fprint(t.out, "\x1b[?1049l")
}


// CSI enable alternate screen buffer
func (t *Term) EnableAlternate() {
    fmt.Fprint(t.out, "\x1b[?1049h")
}


// Flush Term buffer
func (t *Term) Flush() {
    t.out.Flush()
}

