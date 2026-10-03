package kyanbasu

// This package kyanbasu is a terminal ui driver that uses CSI escape codes to manipulate the terminal.
// It has capabilities to set raw mode, move and hide the cursor, enable alternate screens, etc.


import (
    "bufio"
    "fmt"
    "os"
    "golang.org/x/term"
)


type Term struct {
    fd      int
    in      *os.File
    state   *term.State
    out     *bufio.Writer
}


// Initialize a new virtual terminal
func NewTerm() *Term {
    fd := int(os.Stdin.Fd())

    oldState, err := term.MakeRaw(fd)
    if err != nil {
        panic(err)
    }

    t := &Term{fd: fd, in: os.Stdin, state: oldState, out: bufio.NewWriter(os.Stdout),}

    t.EnableAlternate()
    t.HideCursor()
    t.Clear()
    t.Flush()    

    return t
}


// ANSI Clear screen
func (t *Term) Clear() {
    t.out.WriteString("\x1b[2J\x1b[H")
}


// Close virtual terminal and resets to old state
func (t *Term) CloseTerm() {
    term.Restore(t.fd, t.state)
    t.DisableAlternate()
    t.SetColor(string(Reset))
    t.ShowCursor()
    t.Flush()
}


// CSI disable alternate screen buffer
func (t *Term) DisableAlternate() {
    t.out.WriteString("\x1b[?1049l")
}


// CSI enable alternate screen buffer
func (t *Term) EnableAlternate() {
    t.out.WriteString("\x1b[?1049h")
}


// Flush Term buffer
func (t *Term) Flush() {
    t.out.Flush()
}


// Get terminal size
func (t *Term) GetSize() (int, int) {
    w, h, _ := term.GetSize(t.fd)
    return w, h
}


// Hide the terminal cursor
func (t *Term) HideCursor() {
    t.out.WriteString("\x1b[?25l")
}


// Move the terminal cursor
func (t *Term) Move(x, y int) {
    fmt.Fprintf(t.out, "\x1b[%d;%dH", y+1, x+1)
}


// Read key input
func (t *Term) ReadKey() byte {
    b := make([]byte, 1)
    _, err := os.Stdin.Read(b)
    if err != nil {
        panic(err)
    }    

    return b[0]
}


// Set Terminal Color State for printing
func (t *Term) SetColor(color string) {
    fmt.Fprint(t.out, color)
}


// Show the terminal cursor
func (t *Term) ShowCursor() {
    t.out.WriteString("\x1b[?25h")
}

// Print a byte at cursor position
func (t *Term) WriteByte(char byte) error {
    return t.out.WriteByte(char)
}
